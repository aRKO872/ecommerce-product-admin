package client

import (
	"context"
	"fmt"

	pb "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/common"
	"golang.org/x/sync/errgroup"
)


func (c *clientImpl) Heartbeat() (string, error) {
	hg := new(errgroup.Group)
	respChan := make(chan string, 3)
	hg.Go(func() error {
		resp, err := c.inventoryMscClient.Heartbeat(
			context.Background(), 
			&pb.GRPCRequest{},
		)
		if err != nil {
			return err
		}

		respChan <- string(resp.Body)
		return nil
	})

	hg.Go(func() error {
		resp, err := c.productsMscCLient.Heartbeat(
			context.Background(), 
			&pb.GRPCRequest{},
		)
		if err != nil {
			return err
		}

		respChan <- string(resp.Body)
		return nil
	})

	hg.Go(func() error {
		resp, err := c.ordersMscClient.Heartbeat(
			context.Background(), 
			&pb.GRPCRequest{},
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