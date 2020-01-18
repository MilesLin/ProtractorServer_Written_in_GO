package controller

var (
	loginController		 login
)

func Startup() {
	loginController.registerRoutes()
}
