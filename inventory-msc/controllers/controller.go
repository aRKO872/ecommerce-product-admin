package controllers

import (
	invStub "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/inventory-msc"
	"github.com/aRKO872/ecommerce-product-admin/inventory-msc/services"
)

type Controller struct {
	srv services.Service
	invStub.UnimplementedInventoryServiceServer
}

func NewController(srv services.Service) *Controller {
	return &Controller{
		srv: srv,
		UnimplementedInventoryServiceServer: invStub.UnimplementedInventoryServiceServer{},
	}
}