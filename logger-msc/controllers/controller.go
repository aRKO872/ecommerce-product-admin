// Controllers package is responsible for handling HTTP requests and responses.
// It defines the Controller struct, which contains methods to manage various endpoints,
// including health checks and other service-related functionalities.
package controllers

import (
	"encoding/json"
	"log"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/routers"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// Controller struct holds the service layer to process business logic for incoming requests.
type Controller struct {
}

// NewController creates a new instance of Controller with the provided service layer.
func NewController() *Controller {
	return &Controller{}
}

func HandleLogPosting(ev *kafka.Message) error {
	var logMsg routers.LogMessage
	if err := json.Unmarshal(ev.Value, &logMsg); err != nil {
		return err
	}

	log.Printf(LogTemplate, logMsg.AppID, logMsg.CorrelationID, logMsg.Level, logMsg.Message, logMsg.Timestamp)
	return nil
}