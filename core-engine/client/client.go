package client

import (
	"context"

	invStub "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/inventory-msc"
	orderStub "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/orders-msc"
	productStub "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/products-msc"
	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/routers"
)

type clientImpl struct {
	inventoryMscClient invStub.InventoryServiceClient
	productsMscCLient productStub.ProductsServiceClient
	ordersMscClient orderStub.OrdersServiceClient
	logger *routers.Logger
}

func NewClient(
	logger *routers.Logger,
	inventoryMscClient invStub.InventoryServiceClient,
	productsMscCLient productStub.ProductsServiceClient,
	ordersMscClient orderStub.OrdersServiceClient,
) Client {
	return &clientImpl{
		inventoryMscClient: inventoryMscClient,
		productsMscCLient: productsMscCLient,
		ordersMscClient: ordersMscClient,
		logger: logger,
	}
}

type Client interface {
	Heartbeat(ctx context.Context) (string, error)
}