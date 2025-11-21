package services

import "context"

// Heartbeat invokes the client’s Heartbeat method to
// perform a health check.
func (s *serviceImpl) Heartbeat(ctx context.Context) (string, error) {
	return s.client.Heartbeat(ctx)
}