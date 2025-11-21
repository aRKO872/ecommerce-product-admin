package services

import (
	"context"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/routers"
	"github.com/aRKO872/ecommerce-product-admin/core-engine/client"
)

type serviceImpl struct {
	client client.Client
	kafkaProducer *routers.KafkaProducer
	logger *routers.Logger
}

func NewService(logger *routers.Logger, client client.Client, kafkaProducer *routers.KafkaProducer) Service {
	return &serviceImpl{
		logger: logger,
		client: client,
		kafkaProducer: kafkaProducer,
	}
}

type Service interface {
	Heartbeat(ctx context.Context) (string, error)
}