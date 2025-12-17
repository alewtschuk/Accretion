package common

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/text/unicode/norm"
)

func TestNewMnemonic(t *testing.T) {
	language := "english"
	mnemonic := NewMnemonic(language)
	fmt.Printf("Mnemonic WordList: ")
	for _, word := range mnemonic.WordList[language] {
		fmt.Println(word)
	}

	fmt.Printf("Mnemonic Language: %v\n", mnemonic.Language)
	fmt.Printf("Mnemonic WordList Length: %d\n", len(mnemonic.WordList[language]))

	if language == "japanese" {
		fmt.Printf("Mnemonic Delimiter: Ideographic Space. Japanese must be joined by ideographic space")
	} else {
		fmt.Printf("Mnemonic Delimiter: No mnemonic delimiter. Language is not japanese")
	}
}

func TestAvailableLanguages(t *testing.T) {
	language := "english"
	mnemonic := NewMnemonic(language)
	langauges := mnemonic.listLanguages()
	fmt.Printf("Mnemonic languages: %v", langauges)
}

func TestWordNorm(t *testing.T) {
	language := "japanese"
	mnemonic := NewMnemonic(language)
	for _, word := range mnemonic.WordList[language] {
		normWord := mnemonic.normalizeString(word)
		fmt.Printf("Word is: %v\n", word)
		fmt.Printf("Normalized word is: %v\n\n", normWord)
		if !norm.NFKD.IsNormal([]byte(normWord)) {
			t.Fatalf("Word %s is not normalized", normWord)
		}
	}
}

func TestWords(t *testing.T) {
	entries, _ := folder.ReadDir("wordlist")
	for _, entry := range entries {
		language := strings.TrimSuffix(entry.Name(), ".txt")
		mnemonic := NewMnemonic(language)
		fmt.Printf("\033[32mAvailable %s mnemonic words: \x1b[0m%v\n\n", language, mnemonic.WordList[language])
	}
}

func TestLanguageDetection_Unambiguous(t *testing.T) {
	mnemonic := NewMnemonic("english") // Just need one instance to call the method

	// "zone" only exists in English, not in French or other languages
	input := "abandon zone"
	result := mnemonic.detectLanguage(input)

	if result != "english" {
		t.Errorf("Expected 'english', got '%s'", result)
	}
}

func TestLanguageDetection_Prefix(t *testing.T) {
	mnemonic := NewMnemonic("english")

	// "aba" is a prefix of "abandon", "zon" is a prefix of "zone"
	input := "aba zon"
	result := mnemonic.detectLanguage(input)

	if result != "english" {
		t.Errorf("Expected 'english', got '%s'", result)
	}
}

func TestLanguageDetection_Disambiguation(t *testing.T) {
	mnemonic := NewMnemonic("english")

	// "abandon" exists in both English and French
	// "about" exists in English but not French (French has "aboutir")
	input := "abandon about"
	result := mnemonic.detectLanguage(input)

	if result != "english" {
		t.Errorf("Expected 'english', got '%s'", result)
	}
}

func TestLanguageDetection_Spanish(t *testing.T) {
	mnemonic := NewMnemonic("spanish")

	// Words that clearly exist in Spanish
	input := "ábaco abdomen"
	result := mnemonic.detectLanguage(input)

	if result != "spanish" {
		t.Errorf("Expected 'spanish', got '%s'", result)
	}
}

func TestLanguageDetection_Japanese(t *testing.T) {
	mnemonic := NewMnemonic("japanese")

	// Words that clearly exist in Spanish
	input := "りゅうがく まつり"
	result := mnemonic.detectLanguage(input)

	if result != "japanese" {
		t.Errorf("Expected 'japanese', got '%s'", result)
	}
}

func TestLanguageDetection_InvalidWord(t *testing.T) {
	mnemonic := NewMnemonic("english")

	// "xyzabc" doesn't exist in any language's wordlist
	input := "abandon xyzabc"

	// This should panic/fatal, so we need to catch it
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected function to panic on invalid word, but it didn't")
		}
	}()

	mnemonic.detectLanguage(input)
}

func TestLanguageDetection(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		shouldErr bool
	}{
		{
			name:      "unambiguous english",
			input:     "abandon zone",
			expected:  "english",
			shouldErr: false,
		},
		{
			name:      "prefix matching",
			input:     "aba zon",
			expected:  "english",
			shouldErr: false,
		},
		{
			name:      "disambiguation needed",
			input:     "abandon about",
			expected:  "english",
			shouldErr: false,
		},
		{
			name:      "invalid word",
			input:     "abandon xyzabc",
			expected:  "",
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mnemonic := NewMnemonic("english")

			if tt.shouldErr {
				defer func() {
					if r := recover(); r == nil {
						t.Error("Expected function to panic, but it didn't")
					}
				}()
			}

			result := mnemonic.detectLanguage(tt.input)

			if !tt.shouldErr && result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestToEntropy(t *testing.T) {
	// Generate 1024 random 32-byte sequences
	data := make([][]byte, 1024)
	for i := 0; i < 1024; i++ {
		randomBytes := make([]byte, 32)
		rand.Read(randomBytes) // Fill with random data
		data[i] = randomBytes
	}

	// Add the specific test case
	data = append(data, []byte("Lorem ipsum dolor sit amet amet."))

	// Create Mnemonic instance
	mnemonic := NewMnemonic("english")

	// Test each data sequence
	for i, d := range data {
		// Convert to mnemonic
		mnemonicString := mnemonic.toMnemonic(d)

		// Split into words
		words := strings.Split(mnemonicString, " ")

		// Convert back to entropy
		result := mnemonic.toEntropy(words)

		// Check if result equals original data
		if !bytes.Equal(result, d) {
			t.Errorf("Test case %d failed: entropy round-trip mismatch", i)
			t.Errorf("Original: %x", d)
			t.Errorf("Result:   %x", result)
		}
	}
}
