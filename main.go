package main

import (
	"ProtractorServer/controller"
	"net/http"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8080", nil)
}
