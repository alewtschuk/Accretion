// package declaration
package validation

// imports
import (
	"accretion/internal/common/client"
	"accretion/internal/common/genproto/validator"
	"fmt"
	"os"
)

// structure definitions
type Peer struct {
	Hostname string
	conn     validator.ValidatorServiceClient
}

// function definitions
func NewPeer(Hostname string) (*Peer, error) {

	// initialized data
	var err error = nil
	var peer *Peer = nil

	// allocate a peer
	peer = new(Peer)

	// populate fields
	peer.Hostname = Hostname
	peer.conn, _, err = client.NewPeerClient(Hostname + ":3000")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	}

	// done
	return peer, err
}

// method definitions
func (p Peer) Name() string {

	// done
	return p.Hostname
}
