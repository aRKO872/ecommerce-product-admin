package client

import (
	invStub "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/inventory-msc"
	orderStub "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/orders-msc"
	productStub "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/products-msc"
)

type clientImpl struct {
	inventoryMscClient invStub.InventoryServiceClient
	productsMscCLient productStub.ProductsServiceClient
	ordersMscClient orderStub.OrdersServiceClient
}

func NewClient(
	inventoryMscClient invStub.InventoryServiceClient,
	productsMscCLient productStub.ProductsServiceClient,
	ordersMscClient orderStub.OrdersServiceClient,
) Client {
	return &clientImpl{
		inventoryMscClient: inventoryMscClient,
		productsMscCLient: productsMscCLient,
		ordersMscClient: ordersMscClient,
	}
}

type Client interface {
	Heartbeat() (string, error)
}