package controllers

import (
	"net/http"

	utils "github.com/aRKO872/ecommerce-product-admin-microservice-utils/routers"
)

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
