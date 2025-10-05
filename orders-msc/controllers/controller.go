package controllers

import (
	orderStub "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/orders-msc"
	"github.com/aRKO872/ecommerce-product-admin/orders-msc/services"
)

type Controller struct {
	srv services.Service
	orderStub.UnimplementedOrdersServiceServer
}

func NewController(srv services.Service) *Controller {
	return &Controller{
		srv: srv,
		UnimplementedOrdersServiceServer: orderStub.UnimplementedOrdersServiceServer{},
	}
}