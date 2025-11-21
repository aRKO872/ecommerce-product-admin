package prepare

import (
	"log"

	pb "github.com/aRKO872/ecommerce-product-admin-microservice-utils/grpc/inventory-msc"
	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/routers"
	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/utils"
	"github.com/aRKO872/ecommerce-product-admin/inventory-msc/controllers"
	"github.com/aRKO872/ecommerce-product-admin/inventory-msc/models"
	"github.com/aRKO872/ecommerce-product-admin/inventory-msc/services"
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

	srv := services.NewService(logger)
	r := controllers.NewController(srv)

	sr := routers.ServiceRouter{
		AppID:        config.AppID,
	}

	sr.ServeGRPC(func(s *grpc.Server) {
		pb.RegisterInventoryServiceServer(s, r)
	}, r.GetInterceptors()...)

	log.Printf("grpc service %s started", config.AppID)
	sr.Wait()
}