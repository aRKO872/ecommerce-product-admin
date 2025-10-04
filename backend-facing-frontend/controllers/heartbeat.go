package controllers

import "net/http"

func (c *Controller) Heartbeat(rw http.ResponseWriter, r *http.Request) {
	var err error
	status, err := c.srv.Heartbeat()

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(status))
}