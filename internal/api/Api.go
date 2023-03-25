package api

import (
	"log"
	"net/http"
)

type Api struct {
	Server http.Server
	Logger log.Logger
}

type Credential struct {
	Host string
	Port int
}

func New(credential Credential) (*Api, error) {
	// todo:: generate new server
	return &Api{}, nil
}
