// package declaration
package main

// imports
import (
	"accretion/internal/common"
	"fmt"
	"os"
)

// entry point
func main() {
	w, _ := common.NewWallet(os.Args[1])

	fmt.Printf("%+v\n", w)
}
