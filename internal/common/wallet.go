// package declaration
package common

// imports
import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// structure definitions
type Wallet struct {
	Name    string             `json:"name,omitempty"`
	Public  ed25519.PublicKey  `json:"public"`
	Private ed25519.PrivateKey `json:"private"`
}

// construct a wallet
func NewWallet() (*Wallet, error) {

	// initialized data
	var err error = nil
	var wallet *Wallet = new(Wallet)

	// generate a key pair
	wallet.Public, wallet.Private, err = ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}

	// done
	return wallet, nil
}

func (w Wallet) String() string {
	return fmt.Sprintf("%s (%s)", w.Name, string(w.Public))
}

func (w Wallet) MarshalJSON() ([]byte, error) {

	// initialized data
	var publicKey string = base64.StdEncoding.EncodeToString(w.Public)
	var privateKey string = base64.StdEncoding.EncodeToString(w.Private)

	// done
	return json.Marshal(struct {
		Public  string `json:"public"`
		Private string `json:"private"`
	}{
		Public:  publicKey,
		Private: privateKey,
	})
}

func (w *Wallet) UnmarshalJSON(data []byte) error {

	// initialized data
	var err error = nil
	var publicBytes []byte
	var privateBytes []byte
	var transient struct {
		Public  string `json:"public"`
		Private string `json:"private"`
	}

	// unmarshall
	err = json.Unmarshal(data, &transient)
	if err != nil {
		return err
	}

	// decode public key
	publicBytes, err = base64.StdEncoding.DecodeString(transient.Public)
	if err != nil {
		return err
	}

	// decode private key
	privateBytes, err = base64.StdEncoding.DecodeString(transient.Private)
	if err != nil {
		return err
	}

	// store the public and private keys
	w.Public = ed25519.PublicKey(publicBytes)
	w.Private = ed25519.PrivateKey(privateBytes)

	// done
	return nil
}
