// package declaration
package common

// imports
import (
	"crypto/ed25519"
)

// interface definitions
type IWallet interface{}

// structure definitions
type Wallet struct {
	Public  ed25519.PublicKey
	Private ed25519.PrivateKey
}

// construct a wallet from a path
func NewWallet(path string) (
	wallet *Wallet,
	err error,
) {

	// initialized data
	err = nil
	wallet = new(Wallet)

	// done
	return wallet, err
}
