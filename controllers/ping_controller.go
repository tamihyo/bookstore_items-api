package controllers

import "net/http"

const (
	pong = "pong"
)

type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}
type pingController struct{}

var (
	PingController pingControllerInterface = &pingController{}
)

func (s *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
