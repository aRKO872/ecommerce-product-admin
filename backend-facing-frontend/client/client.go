// client package is responsible for communication with external gRPC services.
// It defines the Client interface and its implementation.
package client

import (
	"context"

	pb "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/core-engine"
	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/routers"
)

// clientImpl is the concrete implementation of the Client interface.
//
// It communicates with the Core Engine microservice through a gRPC client
// and delegates heartbeat requests to the underlying CoreEngineServiceClient.
type clientImpl struct {
	logger *routers.Logger
	coreEngineClient pb.CoreEngineServiceClient
}

// NewClient returns a new instance of Client using the provided
// CoreEngineServiceClient.
func NewClient(logger *routers.Logger, coreEngineClient pb.CoreEngineServiceClient) Client {
	return &clientImpl{
		logger: logger,
		coreEngineClient: coreEngineClient,
	}
}

// Client defines the interface for interacting with external services.
type Client interface {
	// Heartbeat performs a health check by invoking the underlying
	// CoreEngineServiceClient's Heartbeat method and returning its result.
	Heartbeat(ctx context.Context) (string, error)
}