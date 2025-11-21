package models

type HeartbeatResponse struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}