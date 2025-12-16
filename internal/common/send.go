// package declaration
package common

// imports
import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
)

// type definitions
type TxSend struct {
	To       ed25519.PublicKey
	From     ed25519.PublicKey
	Quantity uint64
}

// function definitions
func NewTxSend(
	To ed25519.PublicKey,
	From ed25519.PublicKey,
	Quantity uint64,
) *TxSend {

	// done
	return &TxSend{
		To:       To,
		From:     From,
		Quantity: Quantity,
	}
}

// method definitions
func (t TxSend) MarshalJSON() ([]byte, error) {

	// initialized data
	var toKey string = base64.StdEncoding.EncodeToString(t.To)
	var fromKey string = base64.StdEncoding.EncodeToString(t.From)

	// done
	return json.Marshal(struct {
		To       string `json:"to"`
		From     string `json:"from"`
		Quantity uint64 `json:"quantity"`
	}{
		To:       toKey,
		From:     fromKey,
		Quantity: t.Quantity,
	})
}
