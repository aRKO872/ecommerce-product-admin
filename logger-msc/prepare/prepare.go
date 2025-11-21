package prepare

import (
	"errors"
	"log"

	"github.com/aRKO872/ecommerce-product-admin-microservice-utils/utils"
	"github.com/aRKO872/ecommerce-product-admin/logger-msc/controllers"
	"github.com/aRKO872/ecommerce-product-admin/logger-msc/literals"
	"github.com/aRKO872/ecommerce-product-admin/logger-msc/models"
	"github.com/go-playground/validator/v10"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
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

	router := controllers.NewController().Router(config.AppID)

	if err := router.InitKafkaConsumer(LoggerConsumeListener); err != nil {
		log.Fatal("failed to initialize kafka pubsub: ", err.Error())
	}

	router.ServeHTTP()
}

func LoggerConsumeListener(ev *kafka.Message) error {
	key := string(ev.Key)
	switch key {
	case literals.TopicLogPosting:
		return controllers.HandleLogPosting(ev)
	default:
		log.Printf("unknown topic key: %s", key)
		return errors.New("unknown key set in message: "+key)
	}
}