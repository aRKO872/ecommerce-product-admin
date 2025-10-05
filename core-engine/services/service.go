package services

import "github.com/aRKO872/ecommerce-product-admin/core-engine/client"

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