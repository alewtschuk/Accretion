// package declaration
package client

// imports
import (
	"accretion/internal/common/genproto/validator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// function definitions
func NewPeerClient(grpcAddr string) (validator.ValidatorServiceClient, func() error, error) {

	// initialized data
	var err error = nil
	var conn *grpc.ClientConn

	// construct a new grpc client
	conn, err = grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, func() error { return nil }, err
	}

	// done
	return validator.NewValidatorServiceClient(conn), conn.Close, nil
}
