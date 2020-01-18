package controller

import (
	"net/http"
	)

type login struct {

}

func (lo login) registerRoutes() {
	http.HandleFunc("/api/login", lo.login)
	http.HandleFunc("/api/login/currentIdentity", lo.GetCurrentIdentity)
}

func (lo login) login(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello"))
}

func (lo login) GetCurrentIdentity(writer http.ResponseWriter, request *http.Request) {
	
}

