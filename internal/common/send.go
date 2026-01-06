// package declaration
package common

// imports
import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// interfac definitions
type ITxSend interface {
	ITx
}

// structure definitions
type TxSend struct {

	// super
	Tx

	// fields
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
		Tx:       Tx{Signature: nil},
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

func (t TxSend) Sign(w *Wallet) error {

	// initialized data
	var err error = nil
	var data []byte = []byte{}

	// serialize the transaction
	data, err = t.MarshalJSON()
	if err != nil {
		return err
	}

	// sign the transaction
	t.Signature = ed25519.Sign(w.Private, data)

	// done
	return nil
}

func (t TxSend) Verify() bool {

	// initialized data
	var err error = nil
	var data []byte = []byte{}

	// serialize the transaction
	data, err = t.MarshalJSON()
	if err != nil {
		return false
	}

	// done
	return ed25519.Verify(
		t.From,
		data,
		t.Signature,
	)
}

func (t TxSend) Enforce() bool {

	// STUB
	return false
}

func (t TxSend) String() string {

	// done
	return fmt.Sprintf(
		"\033[1;31m0x%s\033[22m sent \033[1;31m%d\033[22m to \033[1;31m0x%s\033[22;0m -> signature: \033[1;34m0x%s\033[22;0m",

		hex.EncodeToString(t.From)[0:6],
		t.Quantity,
		hex.EncodeToString(t.To)[0:6],

		// super
		t.Tx.String()[:8],
	)
}
