// package declaration
package main

// imports
import (
	"accretion/internal/common/genproto/validator"
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
	peers          []*Peer
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
		var p *Peer = nil

		// construct a new peer
		p, err = NewPeer(peer_hostnames[i])
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
	var grpcServer *grpc.Server = nil
	var listen net.Listener = nil
	var service GrpcServer = GrpcServer{
		Peers: peers,
	}

	// construct a grpc server
	grpcServer = grpc.NewServer()

	// register the validator service
	validator.RegisterValidatorServiceServer(grpcServer, service)

	// construct a listener
	listen, err = net.Listen("tcp", ":3000")
	if err != nil {
		os.Exit(1)
	}

	// logs
	fmt.Printf("%s up!\n", host)

	// listen and serve
	grpcServer.Serve(listen)
}
