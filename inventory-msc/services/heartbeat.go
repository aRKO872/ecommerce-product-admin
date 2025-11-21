package services

import (
	"context"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/literals"
)

func (s *serviceImpl) Heartbeat(ctx context.Context) (string) {
	s.logger.Log(ctx, "Heartbeat received", literals.LogLevelInfo)
	return "OK"
}