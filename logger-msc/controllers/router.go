package controllers

import (
	"net/http"

	utils "github.com/aRKO872/ecommerce-product-admin-microservice-utils/routers"
)

// Router configures and returns a ServiceRouter instance for the controller.
//
// It registers all the HTTP routes handled by this controller, mapping each
// endpoint to its corresponding handler function. The router is initialized
// with the provided application identifier and used to integrate controller
// routes into the broader application routing system.
//
// Parameters:
//   - appID: A unique identifier for the application or service to which
//     this router belongs.
//
// Returns:
//   - *utils.ServiceRouter: A configured router containing all routes
//     exposed by this controller.
func (c *Controller) Router(appID string) *utils.ServiceRouter {
	return &utils.ServiceRouter{
		AppID: appID,
		Routes: []utils.CommonRouter{
			{
				Method:   http.MethodGet,
				Endpoint: HealthPath,
				Handler:  c.Heartbeat,
			},
		},
	}
}
