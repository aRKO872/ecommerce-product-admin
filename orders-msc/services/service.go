package services

type serviceImpl struct {
}

func NewService() Service {
	return &serviceImpl{}
}

type Service interface {
	Heartbeat() (string)
}