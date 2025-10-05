package controllers

import (
	productStub "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/products-msc"
	"github.com/aRKO872/ecommerce-product-admin/products-msc/services"
)

type Controller struct {
	srv services.Service
	productStub.UnimplementedProductsServiceServer
}

func NewController(srv services.Service) *Controller {
	return &Controller{
		srv: srv,
		UnimplementedProductsServiceServer: productStub.UnimplementedProductsServiceServer{},
	}
}