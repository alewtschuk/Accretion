// package declaration
package common

// imports
import (
	"encoding/hex"
)

// interface definitions
type ITx interface {

	// Sign a transaction with a wallet
	Sign(w *Wallet) error

	// Verify a transaction's digital signature
	Verify() bool

	// Enforces that transaction fits protocol parameters
	Enforce() bool

	// Serialize a transaction into json
	MarshalJSON() ([]byte, error)

	// Serialize a transaction into a string
	String() string
}

// structure definitions
type Tx struct {
	Signature []byte
}

// method definitions
func (t Tx) String() string {

	// encode the signature as a hexadecimal
	return hex.EncodeToString(t.Signature)
}

/*
const (
	TYPESEND = iota
	TYPEMINT
	TYPEBURN
	TYPESTORE
)

const MAX_TX_SIZE uint16 = 1232

// Transaction definiton
type Tx struct {
	Data      TxData
	Signature TxSig
	//Size      []byte // size of the transaction in bytes
}

// Definition of data held within a transaction
type TxData struct {
	InvolvedAddresses TxAddresses // all addresses involved in a transaction
	Instruction       TxInstruct  // instructions for transaction type
	Hash              []byte      // hash of the transaction
	Time              time.Time   // time that transaction was seen
}

// Defines fields based on transaction instruction
type TxInstruct struct {
	Type        uint8
	Quantity    uint64
	MessageData []byte
}

// Transaction signature definition
type TxSig []byte

// Transaction addresses definition
type TxAddresses []ed25519.PublicKey
type TxInstruction interface {
	encodeInstructions() []byte
}

var transactionTypes = map[int]string{
	TYPESEND: "Send",
	//TypeMint: "Mint",
	//TypeBurn: "Burn",
	//TypeStore: "Store",
}

func (tx *Tx) Sign(w *Wallet) []byte {
	signature := ed25519.Sign(w.Private, tx.Data.Instruction.encodeInstruction())
	return signature
}

// Encode the Instructions to binary to prepare for signing
func (txi *TxInstruct) encodeInstruction() []byte {
	smallBuf := make([]byte, binary.MaxVarintLen64)
	encodedInstruction := binary.PutUvarint(smallBuf, uint64(txi.Type))
	encodedInstruction += binary.PutUvarint(smallBuf[encodedInstruction:], txi.Quantity)

	largeBuf := new(bytes.Buffer)
	lenBuf, err := largeBuf.Write(smallBuf)
	if err != nil || lenBuf > 1232 {
		log.Panicln("ERROR: Buffer too large or instructions over 1232 bytes")
		return nil
	}

	messageLen := int64(len(txi.MessageData))
	err = binary.Write(largeBuf, binary.BigEndian, messageLen)
	if err != nil {
		log.Println("ERROR: Transaction Message Data Length could not be written to buffer")
		return nil
	}

	err = binary.Write(largeBuf, binary.BigEndian, txi.MessageData)
	if err != nil {
		log.Println("ERROR: Transaction Message Data could not be written to buffer")
		return nil
	}

	return largeBuf.Bytes()

}

*/
