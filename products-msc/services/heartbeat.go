package services

import "log"

func (s *serviceImpl) Heartbeat() (string) {
	log.Println("Heartbeat received")
	return "OK"
}