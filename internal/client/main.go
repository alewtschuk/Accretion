// package declaration
package main

// imports
import (
	"accretion/internal/common"
	"fmt"
)

// entry point
func main() {
	w, _ := common.NewWallet()
	x, _ := common.NewWallet()
	t := common.NewTxSend(x.Public, w.Public, 100)
	t.Sign(w)
	d, _ := t.MarshalJSON()

	fmt.Printf("%+v\n", string(d))
}
