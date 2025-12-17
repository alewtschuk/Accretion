// package declaration
package main

// imports
import (
	"accretion/internal/common/genproto/validator"
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

// struct definitions
type GrpcServer struct {
	Peers []*Peer
}

func (g GrpcServer) Heartbeat(
	ctx context.Context,
	in *emptypb.Empty,
) (
	*emptypb.Empty,
	error,
) {

	// log
	fmt.Printf("Ping\n")

	// done
	return nil, nil
}

func (g GrpcServer) Gossip(
	ctx context.Context,
	in *validator.Topic,
) (
	*emptypb.Empty,
	error,
) {

	// iterate through each peer
	for i := 0; i < len(g.Peers); i++ {

		// FIX: causes infinite loop in services
		// _, err := g.Peers[i].conn.Gossip(ctx, in)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		// }
		fmt.Printf("%s, did you know that \"%s\"?\n", g.Peers[i].Name(), in.GetName())
	}

	// done
	return nil, nil
}
