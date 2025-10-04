package client

import pb "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/core-engine"

type clientImpl struct {
	coreEngineClient pb.CoreEngineServiceClient
}

func NewClient(coreEngineClient pb.CoreEngineServiceClient) Client {
	return &clientImpl{
		coreEngineClient: coreEngineClient,
	}
}

type Client interface {
	Heartbeat() (string, error)
}