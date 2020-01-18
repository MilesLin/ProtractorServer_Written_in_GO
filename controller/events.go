package controller

import (
	"ProtractorServer/model"
	"encoding/json"
	"net/http"
)

type events struct {}

func (e events) registerRoutes() {
	http.HandleFunc("/api/events", e.Get)
}

func (e events) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		result, _ := json.Marshal(model.MyEvents)
		w.Write(result)
	}else{
		http.NotFound(w, r)
	}
}
