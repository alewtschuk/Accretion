// package declaration
package common

// imports
import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
)

// type definitions
type TxBurn struct {
	From     ed25519.PublicKey
	Quantity uint64
}

// function definitions
func NewTxBurn(
	From ed25519.PublicKey,
	Quantity uint64,
) *TxBurn {

	// done
	return &TxBurn{
		From:     From,
		Quantity: Quantity,
	}
}

// method definitons
func (t TxBurn) MarshalJSON() ([]byte, error) {

	// initialized data
	var fromKey string = base64.StdEncoding.EncodeToString(t.From)

	// done
	return json.Marshal(struct {
		From     string `json:"from"`
		Quantity uint64 `json:"quantity"`
	}{
		From:     fromKey,
		Quantity: t.Quantity,
	})
}
