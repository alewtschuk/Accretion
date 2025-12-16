// package declaration
package common

// imports
import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
)

// type definitions
type TxFaucet struct {
	To       ed25519.PublicKey
	Quantity uint64
}

// function definitions
func NewTxFaucet(
	To ed25519.PublicKey,
	Quantity uint64,
) *TxFaucet {

	// done
	return &TxFaucet{
		To:       To,
		Quantity: Quantity,
	}
}

// method definitions
func (t TxFaucet) MarshalJSON() ([]byte, error) {

	// initialized data
	var toKey string = base64.StdEncoding.EncodeToString(t.To)

	// done
	return json.Marshal(struct {
		To       string `json:"to"`
		Quantity uint64 `json:"quantity"`
	}{
		To:       toKey,
		Quantity: t.Quantity,
	})
}
