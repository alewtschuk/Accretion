// package declaration
package common

// imports
import (
	"crypto/ed25519"
	"encoding/base64"
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
		"%s sent %d to %s; signature: %s",

		base64.StdEncoding.EncodeToString(t.From),
		t.Quantity,
		base64.StdEncoding.EncodeToString(t.To),

		// super
		t.Tx.String(),
	)
}
