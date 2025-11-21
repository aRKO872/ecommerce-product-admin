package prepare

import (
	"log"

	pb "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/core-engine"
	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/routers"
	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/utils"
	"github.com/aRKO872/ecommerce-product-admin/core-engine/client"
	"github.com/aRKO872/ecommerce-product-admin/core-engine/controllers"
	"github.com/aRKO872/ecommerce-product-admin/core-engine/models"
	"github.com/aRKO872/ecommerce-product-admin/core-engine/services"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
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

	kafkaProducer, err := routers.NewKafkaProducer()
	if err != nil {
		log.Fatal("failed to create kafka producer: ", err.Error())
	}

	logger := routers.NewLogger(kafkaProducer, config.AppID)
	grpcConfig := utils.NewGRPCConfig()
	inventoryMscClient := grpcConfig.GetInventoryMscClient()
	ordersMscClient := grpcConfig.GetOrdersMscClient()
	productsMscClient := grpcConfig.GetProductsMscClient()

	client := client.NewClient(logger, inventoryMscClient, ordersMscClient, productsMscClient)
	
	srv := services.NewService(logger, client, kafkaProducer)
	r := controllers.NewController(srv)

	sr := routers.ServiceRouter{
		AppID:        config.AppID,
		IsKafkaEnabled: true,
		PubsubProducer: kafkaProducer,
	}

	sr.ServeGRPC(func(s *grpc.Server) {
		pb.RegisterCoreEngineServiceServer(s, r)
	}, r.GetInterceptors()...)

	log.Printf("grpc service %s started", config.AppID)
	sr.Wait()
}