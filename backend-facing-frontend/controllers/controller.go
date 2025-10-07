// Controllers package is responsible for handling HTTP requests and responses.
// It defines the Controller struct, which contains methods to manage various endpoints,
// including health checks and other service-related functionalities.
package controllers

import "github.com/aRKO872/ecommerce-product-admin/backend-facing-frontend/services"

// Controller struct holds the service layer to process business logic for incoming requests.
type Controller struct {
	srv services.Service
}

// NewController creates a new instance of Controller with the provided service layer.
func NewController(srv services.Service) *Controller {
	return &Controller{srv: srv}
}