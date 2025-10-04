package controllers

import (
	utils "github.com/aRKO872/ecommerce-product-admin-microservice-utils/routers"
	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc-interceptors"
	"google.golang.org/grpc"
)

func (c *Controller) Router(appID string) *utils.ServiceRouter {
	return &utils.ServiceRouter{
		AppID: appID,
	}
}

func (c *Controller) GetInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{
		grpcinterceptors.ErrorInterceptor(),
	}
}