package prepare

import (
	"log"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/utils"
	"github.com/aRKO872/ecommerce-product-admin/backend-facing-frontend/client"
	"github.com/aRKO872/ecommerce-product-admin/backend-facing-frontend/controllers"
	"github.com/aRKO872/ecommerce-product-admin/backend-facing-frontend/models"
	"github.com/aRKO872/ecommerce-product-admin/backend-facing-frontend/services"
	"github.com/go-playground/validator/v10"
)

func Prepare() {
	var config models.Config

	if err := utils.ParseEnv(&config); err != nil {
		log.Fatal("failed to parse env variables: ", err.Error())
	}

	v := validator.New()
	if err := v.Struct(config); err != nil {
		log.Fatal("env validation failed: ", err.Error())
	}

	grpcConfig := utils.NewGRPCConfig()

	coreEngineClient := grpcConfig.GetCoreEngineClient()
	client := client.NewClient(coreEngineClient)

	svc := services.NewService(client)
	router := controllers.NewController(svc).Router(config.AppID)

	router.ServeHTTP()
}