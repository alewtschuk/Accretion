// package declaration
package main

// imports
import (
	"accretion/internal/common/genproto/mempool"
	"accretion/internal/common/genproto/validator"
	mmp "accretion/internal/validator/mempool"
	"accretion/internal/validator/validation"
	"fmt"
	"net"
	"os"
	"slices"
	"strings"

	"google.golang.org/grpc"
)

// data
var (
	host           string
	peer_hostnames []string
	peers          []*validation.Peer
)

// early
func init() {

	// get our hostname
	host, _ = os.Hostname()

	// get peer hostnames
	peer_hostnames = strings.Split(os.Getenv("PEERS"), ";")

	// remove this node from peers
	idx := slices.Index(peer_hostnames, host)
	if idx != -1 {
		peer_hostnames = slices.Delete(peer_hostnames, idx, idx+1)
	}

	// who are our peers
	for i := 0; i < len(peer_hostnames); i++ {

		// initialized data
		var err error = nil
		var p *validation.Peer = nil

		// construct a new peer
		p, err = validation.NewPeer(peer_hostnames[i])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to connect to %s\n", peer_hostnames[i])
		}

		// append the peer to the peer list
		peers = append(peers, p)
	}
}

// entry point
func main() {

	// initialized data
	var err error = nil
	var grpcMempool *grpc.Server = nil
	var grpcValidator *grpc.Server = nil

	var mplisten net.Listener = nil
	var valisten net.Listener = nil

	var mp mmp.GrpcMemPool = mmp.GrpcMemPool{
		Pending: mmp.NewMemPool(),
	}
	var v GrpcValidator = GrpcValidator{}

	// construct a grpc server
	grpcMempool = grpc.NewServer()
	grpcValidator = grpc.NewServer()

	// register the mempool service
	mempool.RegisterMemPoolServiceServer(grpcMempool, &mp)

	// register the validator service
	validator.RegisterValidatorServiceServer(grpcValidator, &v)

	// construct a listener
	mplisten, err = net.Listen("tcp", ":3000")
	if err != nil {
		os.Exit(1)
	}

	// construct a listener
	valisten, err = net.Listen("tcp", ":3001")
	if err != nil {
		os.Exit(1)
	}

	go func() {
		// listen and serve
		grpcValidator.Serve(valisten)
	}()

	// listen and serve
	grpcMempool.Serve(mplisten)
}
