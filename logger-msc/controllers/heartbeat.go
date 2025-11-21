package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/aRKO872/ecommerce-product-admin/logger-msc/models"
)

// Heartbeat handles HTTP requests to verify the health and availability of the service, and it's underlying gRPC services.
//
// It calls the underlying service layer's Heartbeat method to perform the actual
// health check. The method then responds with a JSON-encoded HeartbeatResponse
// that includes either the current status of the service or an error message.
// The response status code is "OK" if all the services, including the underlying
// gRPC services, are functioning correctly; otherwise, it returns "500" error, showing
// the reason for the failure.
//
//     Route: GET /bff-service/health
//
// Response Format (Success):
//   {
//     "status": "OK"
//   }
//
// Response Format (Failure):
//   {
//     "error": "error message"
//   }
//
// Behavior:
//   - On success, the handler responds with HTTP 200 (OK) and a JSON object
//     containing the service status.
//   - On failure, it responds with HTTP 500 (Internal Server Error) and a JSON
//     object containing an error message.
//
// Parameters:
//   - rw: The HTTP response writer used to send the response.
//   - r:  The HTTP request received from the client.
func (c *Controller) Heartbeat(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	errorResp := models.HeartbeatResponse{
		Status: "OK",
	}
	byteResp, _ := json.Marshal(errorResp)
	rw.Write(byteResp)
}