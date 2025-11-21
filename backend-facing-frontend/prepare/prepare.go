// Package prepare initializes and bundles together the core components
// required for the Backend-Facing-Frontend (BFF) service.
//
// It is responsible for reading and validating environment config, setting up
// gRPC connection to core-engine microservice, and preparing the HTTP
// routing layer for incoming requests. This package serves as the entry point
// for bootstrapping the BFF service runtime.
package prepare

import (
	"log"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/routers"
	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/utils"
	"github.com/aRKO872/ecommerce-product-admin/backend-facing-frontend/client"
	"github.com/aRKO872/ecommerce-product-admin/backend-facing-frontend/controllers"
	"github.com/aRKO872/ecommerce-product-admin/backend-facing-frontend/models"
	"github.com/aRKO872/ecommerce-product-admin/backend-facing-frontend/services"
	"github.com/go-playground/validator/v10"
)

// Prepare initializes the Backend-Facing-Frontend (BFF) service.
//
// It performs the following steps:
//   1. Parses environment variables.
//   2. Validates the loaded config variables.
//   3. Initializes the gRPC configuration and Core Engine client.
//   4. Constructs the service and controller layers.
//   5. Registers and serves all defined HTTP routes.
//
// The function logs fatal errors and terminates the application if any
// initialization step fails.
//
// Dependencies:
//   - utils: For environment parsing and gRPC configuration.
//   - validator: For environment struct validation.
//   - client: For gRPC client creation.
//   - services: For business logic binding.
//   - controllers: For route definition and HTTP server setup.
//
// On successful execution, the BFF service starts serving its defined routes.
func Prepare() {
	var config models.Config

	if err := utils.ParseEnv(&config); err != nil {
		log.Fatal("failed to parse env variables: ", err.Error())
	}

	v := validator.New()
	if err := v.Struct(config); err != nil {
		log.Fatal("env validation failed: ", err.Error())
	}

	kafkaProducer, err := routers.NewKafkaProducer()
	if err != nil {
		log.Fatal("failed to create kafka producer: ", err.Error())
	}

	logger := routers.NewLogger(kafkaProducer, config.AppID)

	grpcConfig := utils.NewGRPCConfig()

	coreEngineClient := grpcConfig.GetCoreEngineClient()
	client := client.NewClient(logger, coreEngineClient)

	svc := services.NewService(logger, client)
	router := controllers.NewController(svc).Router(config.AppID)

	router.ServeHTTP()
}