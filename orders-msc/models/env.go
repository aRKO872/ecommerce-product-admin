package models

type Config struct {
	AppID string `env:"APP_ID" validate:"required"`
}