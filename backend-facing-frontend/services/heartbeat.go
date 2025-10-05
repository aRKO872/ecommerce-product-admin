package services

func (s *serviceImpl) Heartbeat() (string, error) {
	return s.client.Heartbeat()
}