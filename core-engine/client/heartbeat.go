package client

import (
	"context"
	"fmt"

	pb "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/common"
	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/utils"
	"golang.org/x/sync/errgroup"
)


func (c *clientImpl) Heartbeat(ctx context.Context) (string, error) {
	hg := new(errgroup.Group)
	respChan := make(chan string, 3)
	grpcReq := &pb.GRPCRequest{
		Metadata: utils.GetGRPCMetadataWithContext(ctx),
	}
	hg.Go(func() error {
		resp, err := c.inventoryMscClient.Heartbeat(
			ctx, 
			grpcReq,
		)
		if err != nil {
			return err
		}

		respChan <- string(resp.Body)
		return nil
	})

	hg.Go(func() error {
		resp, err := c.productsMscCLient.Heartbeat(
			ctx, 
			grpcReq,
		)
		if err != nil {
			return err
		}

		respChan <- string(resp.Body)
		return nil
	})

	hg.Go(func() error {
		resp, err := c.ordersMscClient.Heartbeat(
			ctx, 
			grpcReq,
		)
		if err != nil {
			return err
		}

		respChan <- string(resp.Body)
		return nil
	})

	if err := hg.Wait(); err != nil {
		return "", err
	}

	close(respChan)
	for resp := range respChan {
		if resp != "OK" {
			return "", fmt.Errorf("grpc heartbeat failed. received unexpected msg: %s", resp)
		}
	}

	return "OK", nil
}