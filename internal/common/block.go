// package declaration
package common

// imports
//

// interface definitions
type IBlock interface {

	// compute the hash of the block
	Hash() []byte

	// get the hash of the previous block
	PrevHash() []byte

	// verify the block
	Verify() error

	// validate the block
	Validate() bool

	// compute the hash of the merkle tree root
	MerkleRoot() []byte

	// serialize the block to json
	MarshalJSON() ([]byte, error)
}

type IBlockBuilder interface {

	// super
	IBlock

	// add a transaction to the block
	AddTransaction(tx ITx) error

	// seal the block
	Build() (IBlock, error)
}

// structure definitions
type BlockHeader struct {
	Version      uint32
	PreviousHash []byte
	MerkleRoot   []byte
	Timestamp    int64
	Nonce        uint64
	Difficulty   uint64
	StateRoot    []byte
}

type Block struct {
	Header BlockHeader
	Txns   []ITx
}

type BlockBuilder struct {
	Block
}

// function definitions
func NewBlockBuilder() IBlockBuilder {

	// initialized data

	return &BlockBuilder{
		Block: Block{
			Header: BlockHeader{},
			Txns:   make([]ITx, 0),
		},
	}
}

func (bb *BlockBuilder) AddTransaction(tx ITx) error {

	// done
	bb.Txns = append(bb.Txns, tx)

	// done
	return nil
}

func (b *BlockBuilder) Build() (IBlock, error) {

	// done
	return &b.Block, nil
}

func (b *Block) Hash() []byte {

	// STUB
	return nil
}

func (b *Block) PrevHash() []byte {

	// STUB
	return b.Header.PreviousHash
}

func (b *Block) Verify() error {

	// STUB
	return nil
}

func (b *Block) Validate() bool {

	// STUB
	return false
}

func (b *Block) MerkleRoot() []byte {

	// STUB
	return nil
}

func (b *Block) MarshalJSON() ([]byte, error) {

	// STUB
	return nil, nil
}
