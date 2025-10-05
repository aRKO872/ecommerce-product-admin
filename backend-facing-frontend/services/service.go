package services

import "github.com/aRKO872/ecommerce-product-admin/backend-facing-frontend/client"

type serviceImpl struct {
	client client.Client
}

func NewService(client client.Client) Service {
	return &serviceImpl{
		client: client,
	}
}

type Service interface {
	Heartbeat() (string, error)
}