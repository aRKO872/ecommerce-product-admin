package models

type Config struct {
	AppID          string `env:"APP_ID" validate:"required"`
	CoreEngineAddr string `env:"CORE_ENGINE_ADDR" validate:"required"`
}