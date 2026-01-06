// package declaration
package mempool

// imports
import (
	"accretion/internal/common"
	"accretion/internal/common/genproto/mempool"
	"context"
	"crypto/ed25519"
	"encoding/base64"

	"google.golang.org/protobuf/types/known/emptypb"
)

// struct definitions
type GrpcMemPool struct {
	Pending *MemPool
}

func (g GrpcMemPool) Gossip(
	ctx context.Context,
	in *mempool.Topic,
) (
	*emptypb.Empty,
	error,
) {

	// STUB
	return nil, nil
}

func (g GrpcMemPool) Indulge(
	ctx context.Context,
	in *emptypb.Empty,
) (
	*mempool.TopicsResponse,
	error,
) {

	// STUB
	return nil, nil
}

func (g GrpcMemPool) Send(
	ctx context.Context,
	in *mempool.TxSend,
) (
	*emptypb.Empty,
	error,
) {

	// initialized data
	var err error = nil
	var Tx *common.TxSend
	var to ed25519.PublicKey
	var from ed25519.PublicKey
	var toBytes []byte
	var fromBytes []byte
	var signature []byte
	var toString string
	var fromString string
	var quantity uint64

	// store the signature
	signature = in.GetSignature()

	// store the quantity of funds
	quantity = in.GetQuantity()

	// store the base64 encoded public keys
	toString = in.GetTo()
	fromString = in.GetFrom()

	// decode the receiver
	toBytes, err = base64.StdEncoding.DecodeString(toString)
	if err != nil {
		return nil, err
	}

	// decode the sendr
	fromBytes, err = base64.StdEncoding.DecodeString(fromString)
	if err != nil {
		return nil, err
	}

	// store the public and private keys
	to = ed25519.PublicKey(toBytes)
	from = ed25519.PublicKey(fromBytes)

	// construct a send transaction
	Tx = common.NewTxSend(to, from, quantity)

	// store the digital signature
	Tx.Signature = signature

	// add the send transaction to the mempool
	g.Pending.Add(Tx)

	// done
	return nil, nil
}
