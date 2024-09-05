// main.go
package main

import (
	"fmt"
	"net/http"
	"coffee-project/controllers"
	"coffee-project/models"
)
func main() {
	eventManager := models.NewEventManager()
	controller := controllers.NewEventController(eventManager)

	http.HandleFunc("/events", controller.ListEventsHandler)
	http.HandleFunc("/event", controller.ViewEventDetailsHandler)
	http.HandleFunc("/register", controller.RegisterUserHandler)

	fmt.Println("Server is starting on port 9090...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}



