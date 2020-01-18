package controller

var (
	loginController		 login
	eventsController	 events
)

func Startup() {
	loginController.registerRoutes()
	eventsController.registerRoutes()
}
