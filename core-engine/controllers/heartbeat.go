package controllers

import (
	"context"

	pb "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/common"
)

func (c *Controller) Heartbeat(ctx context.Context, req *pb.GRPCRequest) (*pb.GRPCResponse, error) {
	status, err := c.srv.Heartbeat()
	if err != nil {
		return nil, err
	}
	return &pb.GRPCResponse{Body: []byte(status)}, nil
}