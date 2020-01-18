package main

import (
	"ProtractorServer/controller"
	"ProtractorServer/middleware"
	"fmt"
	"net/http"
)

func main() {
	controller.Startup()

	fmt.Println("Server is running on localhost:5000")
	http.ListenAndServe(":5000", &middleware.AllowCors{})
}
