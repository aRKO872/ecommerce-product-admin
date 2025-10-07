package client

import (
	"context"

	pb "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/common"
)

// Heartbeat calls the Core Engine gRPC service's Heartbeat method and
// returns its response as a string.
//
// If the gRPC call fails, it returns an error describing the issue.
//
// Example:
//
//   status, err := client.Heartbeat()
//   if err != nil {
//       log.Fatal("heartbeat failed:", err)
//   }
//   fmt.Println("Service status:", status)
func (c *clientImpl) Heartbeat() (string, error) {
	resp, err := c.coreEngineClient.Heartbeat(context.Background(), &pb.GRPCRequest{})
	if err != nil {
		return "", err
	}

	return string(resp.Body), nil
}