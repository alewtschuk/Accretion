// package declaration
package main

// imports
import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

// struct definitions
type GrpcValidator struct{}

func (g GrpcValidator) Heartbeat(
	ctx context.Context,
	in *emptypb.Empty,
) (
	*emptypb.Empty,
	error,
) {

	// log
	fmt.Printf("DING DONG!!!\n")

	// done
	return nil, nil
}
