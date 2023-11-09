package handlers

import (
	"fmt"
	"net/http"
)

type HealthCheckHandler interface {
	Handle(http.ResponseWriter, *http.Request)
}

type healthCheckHandler struct {
}

func NewHealthCheckHandler() *healthCheckHandler {
	return &healthCheckHandler{}
}

func (g healthCheckHandler) Handle(w http.ResponseWriter, r *http.Request) {

	// Check the health of the server and return a status code accordingly
	if serverIsHealthy() {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "NDV SERVER IS RUNNING")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Server is not healthy")
	}

}
func serverIsHealthy() bool {
	// Check the health of the server and return true or false accordingly
	// For example, check if the server can connect to the database
	return true
}
