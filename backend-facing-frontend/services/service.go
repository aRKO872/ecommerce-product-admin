// Service layer is responsible for business logic and communication with external services.
// It defines the Service interface and its implementation.
package services

import (
	"context"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/routers"
	"github.com/aRKO872/ecommerce-product-admin/backend-facing-frontend/client"
)

// serviceImpl is the concrete implementation of the Service interface.
// It acts as a middle layer between the controller and client.
type serviceImpl struct {
	client client.Client
	logger *routers.Logger
}

// NewController creates a new instance of Controller with the provided Client layer.
func NewService(logger *routers.Logger, client client.Client) Service {
	return &serviceImpl{
		logger: logger,
		client: client,
	}
}

// Service defines the interface for business logic.
type Service interface {
	// Heartbeat performs a health check by invoking the underlying client’s
	// heartbeat mechanism and returning its result.
	Heartbeat(ctx context.Context) (string, error)
}