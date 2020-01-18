package main

import (
	"ProtractorServer/controller"
	"fmt"
	"net/http"
)

func main() {
	controller.Startup()
	fmt.Println("Server is running on localhost:5000")
	http.ListenAndServe(":5000", nil)
}
