package services

import "log"

func (s *serviceImpl) Heartbeat() (string, error) {
	log.Println("Heartbeat received")
	resp, err := s.client.Heartbeat()
	if err != nil {
		return "", err
	}
	return resp, nil
}