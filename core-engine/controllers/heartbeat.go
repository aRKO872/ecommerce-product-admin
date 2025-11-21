package controllers

import (
	"context"

	pb "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/common"
	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/utils"
)

func (c *Controller) Heartbeat(ctx context.Context, req *pb.GRPCRequest) (*pb.GRPCResponse, error) {
	ctx = utils.PopulateContextWithGRPCMetadata(ctx, req.GetMetadata())
	status, err := c.srv.Heartbeat(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GRPCResponse{Body: []byte(status)}, nil
}