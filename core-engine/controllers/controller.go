package controllers

import (
	ceStub "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/core-engine"
	"github.com/aRKO872/ecommerce-product-admin/core-engine/services"
)

type Controller struct {
	srv services.Service
	ceStub.UnimplementedCoreEngineServiceServer
}

func NewController(srv services.Service) *Controller {
	return &Controller{
		srv: srv,
		UnimplementedCoreEngineServiceServer: ceStub.UnimplementedCoreEngineServiceServer{},
	}
}