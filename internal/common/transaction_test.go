package common

// import (
// 	"bytes"
// 	"encoding/binary"
// 	"fmt"
// 	"testing"
// )

// var txinst *TxInstruct = &TxInstruct{
// 	Type:        0,
// 	Quantity:    10,
// 	MessageData: []byte("This is a test"),
// }

// func TestEncodeInstructions(t *testing.T) {
// 	encodedInst := txinst.encodeInstruction()
// 	fmt.Printf("Encoded binary data (length %d): %v\n", len(encodedInst), encodedInst)
// }

// func TestDecodeInstructions(t *testing.T) {
// 	encodedInst := txinst.encodeInstruction()

// 	buf := bytes.NewReader(encodedInst)
// 	var typ uint8
// 	var quantity uint64
// 	var messageLen uint64

// 	binary.Read(buf, binary.BigEndian, &typ)
// 	binary.Read(buf, binary.BigEndian, &quantity)
// 	binary.Read(buf, binary.BigEndian, &messageLen)
// 	dataRead := make([]byte, messageLen)
// 	binary.Read(buf, binary.BigEndian, &dataRead)
// }
