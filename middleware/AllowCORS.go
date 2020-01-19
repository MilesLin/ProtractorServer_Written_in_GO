package middleware

import (
	"net/http"
)

type AllowCors struct {
	Next http.Handler
}

func (ac *AllowCors) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if ac.Next == nil {
		ac.Next = http.DefaultServeMux
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	ac.Next.ServeHTTP(w, r)
}
