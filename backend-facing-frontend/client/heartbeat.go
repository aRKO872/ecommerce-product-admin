package client

import (
	"context"

	pb "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/common"
)

func (c *clientImpl) Heartbeat() (string, error) {
	resp, err := c.coreEngineClient.Heartbeat(context.Background(), &pb.GRPCRequest{})
	if err != nil {
		return "", err
	}

	return string(resp.Body), nil
}