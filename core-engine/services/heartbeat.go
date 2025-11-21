package services

import (
	"context"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/literals"
)

func (s *serviceImpl) Heartbeat(ctx context.Context) (string, error) {
	s.logger.Log(ctx, "Heartbeat received", literals.LogLevelInfo)
	resp, err := s.client.Heartbeat(ctx)
	if err != nil {
		return "", err
	}
	return resp, nil
}