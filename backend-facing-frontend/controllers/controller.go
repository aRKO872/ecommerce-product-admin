package controllers

import "github.com/aRKO872/ecommerce-product-admin/backend-facing-frontend/services"

type Controller struct {
	srv services.Service
}

func NewController(srv services.Service) *Controller {
	return &Controller{srv: srv}
}