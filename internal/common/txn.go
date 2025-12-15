package common

import "crypto/ed25519"

type txn struct {
	rcv ed25519.PublicKey
	snd ed25519.PublicKey
	qty int
}
