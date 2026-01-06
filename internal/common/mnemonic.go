package common

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"embed"
	"log"
	"os"
	"slices"
	"strings"

	"golang.org/x/text/unicode/norm"
)

const PBKDF2_ROUNDS = 2048

//go:embed wordlist/*.txt
var folder embed.FS

// Defines the mnemonid
type Mnemonic struct {
	Language  string
	WordList  map[string][]string
	Delimiter string
}

// Build the Mnmonic
func NewMnemonic(language string) *Mnemonic {

	var folderName string = "wordlist"
	var delimiter string
	var wordlist map[string][]string = make(map[string][]string)

	s := []string{folderName, language + ".txt"}
	content, err := folder.ReadFile(strings.Join(s, "/"))
	if err != nil {
		log.Printf("ERROR: Cannot read file %s", strings.Join(s, "/"))
		os.Exit(1)
	}

	contentString := bytes.NewBuffer(content).String() // convert content of file to continuous string
	contentStrings := strings.Split(contentString, "\n")
	if len(contentStrings) != 2048 {
		log.Printf("ERROR: Wordlist %s does not contain the correct number(2048) of words", language)
		os.Exit(1)
	}

	// Japanese must be joined by ideographic space
	if language == "japanese" {
		delimiter = "\u3000"
	}

	// Populate the map with the slice of words for each language
	wordlist[language] = contentStrings
	return &Mnemonic{
		Language:  language,
		WordList:  wordlist,
		Delimiter: delimiter,
	}
}

// Lists all supported languages
func (mn *Mnemonic) listLanguages() []string {
	entries, _ := folder.ReadDir("wordlist")
	var languages []string
	for _, entry := range entries {
		if strings.Contains(entry.Name(), ".txt") {
			languages = append(languages, strings.TrimSuffix(entry.Name(), ".txt"))
		}
	}
	return languages
}

// Normalizes the unicode strings to NFKD format
func (mn *Mnemonic) normalizeString(text string) string {
	return norm.NFKD.String(text)
}

// Filteres and detects the language of the passed in code (mnemonic) word set.
func (mn *Mnemonic) detectLanguage(code string) string {
	code = mn.normalizeString(code)
	languages := mn.listLanguages()

	// Create and populate a slice of all possible mnemonics
	var possible []*Mnemonic
	for _, language := range languages {
		p := NewMnemonic(language)
		possible = append(possible, p)
	}

	// Separate the words in the submited mnemonic and build the set for the submited words
	words := strings.Fields(code)
	wordSet := make(map[string]struct{})
	for _, word := range words {
		wordSet[word] = struct{}{}
	}

	// Check which languages the mnemonic could be by filtering based on prefix
	for word := range wordSet {
		var filtered []*Mnemonic
		for _, p := range possible {
			hasMatch := false
			for _, candidate := range p.WordList[p.Language] {
				if strings.HasPrefix(candidate, word) { // if the prefix of the word appears as a word in the wordSet set it as a possible match
					hasMatch = true
					break
				}
			}
			if hasMatch {
				filtered = append(filtered, p)
			}
		}
		possible = filtered // reduce possible mnemonic languages to what passed the prefix filter

		if len(possible) == 0 {
			log.Fatalf("ERROR: Language unrecognized for '%s'", word)
		}
	}

	// If only one possible language remains the lang is found
	if len(possible) == 1 {
		return possible[0].Language
	}

	// If multiple possible languages remain ensure that the language is found based on exact word occurance
	complete := make(map[string]struct{})
	for word := range wordSet {
		var exact []*Mnemonic
		for _, p := range possible {
			if slices.Contains(p.WordList[p.Language], word) { // if the list of words for the language contains the word you have an exact result
				exact = append(exact, p)
			}
		}
		if len(exact) == 1 {
			complete[exact[0].Language] = struct{}{} // mark completed language as the lang of the exact mnemonic
		}
	}

	if len(complete) == 1 {
		for lang := range complete {
			return lang
		}
	}

	var langNames []string
	for _, p := range possible {
		langNames = append(langNames, p.Language)
	}
	log.Fatalf("Language ambiguous between %s", strings.Join(langNames, ", "))
	return ""
}

/*
Create a new mnemonic using a random generated number as entropy.

	As defined in BIP39, the entropy must be a multiple of 32 bits, and its size must be between 128 and 256 bits.
	Therefore the possible values for `strength` are 128, 160, 192, 224 and 256.

	If not provided, the default entropy length will be set to 128 bits.

	The return is a list of words that encodes the generated entropy.
*/
func (mn *Mnemonic) generate(strength uint16) string {

	if !(strength%32 != 0 && (strength < 128 || strength > 256)) {
		log.Fatalf("ERROR: Invalid strength value. Allowed values are [128, 160, 192, 224, 256]")
	}

	// Create and fill the random byte slice
	b := make([]byte, strength/8)
	rand.Read(b)

	return mn.toMnemonic(b)

}

func (mn *Mnemonic) toEntropy(words []string) []byte {
	if len(words)%3 > 0 {
		log.Fatalf("ERROR: Number of words must be one of the following: [12, 15, 18, 21, 24]")
	}

	if len(words) == 0 {
		log.Fatalf("ERROR: Words list is empty")
	}

	// Look up all the words in the list and construct the
	// concatenation of the original entropy and the checksum.
	concatLenBits := len(words) * 11
	concatBits := make([]bool, concatLenBits)
	wordIndex := 0
	for _, word := range words {
		// Find the words index in the wordlist
		ndx, _ := slices.BinarySearch(mn.WordList[mn.Language], word)
		if ndx < 0 {
			log.Fatalf("ERROR: Unable to find '%s' in word list", word)
		}

		// Set the next 11 bits to the value of the index
		for ii := range 11 {
			concatBits[(wordIndex*11)+ii] = (ndx & (1 << (10 - ii))) != 0
		}
		wordIndex++
	}

	checkSumLengthBits := concatLenBits / 30
	entropyLengthBits := concatLenBits - checkSumLengthBits

	// Extract original entropy as bytes
	entropy := make([]byte, entropyLengthBits/8)
	for ii := range len(entropy) {
		for jj := range 8 {
			if concatBits[(ii*8)+jj] {
				entropy[ii] |= 1 << (7 - jj)
			}
		}
	}

	hash := sha256.New()
	hash.Write(entropy)
	hashBytes := hash.Sum(nil)
	hashBits := make([]bool, 0, len(hashBytes)*8) // create bool slice to act as bit slice
	for _, byt := range hashBytes {
		for i := range 8 {
			hashBits = append(hashBits, (byt&(1<<(7-i))) != 0)
		}
	}

	// Check all the checksum bits
	for i := range checkSumLengthBits {
		if concatBits[entropyLengthBits+i] != hashBits[i] {
			log.Fatalf("ERROR: Failed checksum")
		}
	}

	return entropy
}

func (mn *Mnemonic) toMnemonic([]byte) string {

	// STUB
	return ""
}
