package services

import (
	"context"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/routers"
)

type serviceImpl struct {
	logger *routers.Logger
}

func NewService(logger *routers.Logger) Service {
	return &serviceImpl{
		logger: logger,
	}
}

type Service interface {
	Heartbeat(ctx context.Context) (string)
}