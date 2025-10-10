package services

// Heartbeat invokes the client’s Heartbeat method to
// perform a health check.
func (s *serviceImpl) Heartbeat() (string, error) {
	return s.client.Heartbeat()
}