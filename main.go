package main

import (
	"ProtractorServer/controller"
	"net/http"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":5000", nil)
}
