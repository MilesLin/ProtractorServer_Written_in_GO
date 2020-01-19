package main

import (
	"ProtractorServer/controller"
	"ProtractorServer/middleware"
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	controller.Startup()

	fmt.Println("Server is running on localhost:5000")
	err := http.ListenAndServe(":5000", &middleware.AllowCors{})
	fmt.Println(err)
	fmt.Print("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
