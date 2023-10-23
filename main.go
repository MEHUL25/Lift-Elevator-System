package main

import (
	"fmt"
	"net/http"
	"time"

	handlers "github.com/MEHUL25/lift-management-system/handlers"
	"github.com/MEHUL25/lift-management-system/logic"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	// Create an instance of the lift logic
	liftLogic := logic.NewLiftLogic()

	// Define routes
	r.Post("/api/lift", handlers.HandleLiftRequest(liftLogic))

	// Start a Goroutine to continuously log the current floor
	go func() {
		for {
			fmt.Printf("\n-------------------------\nCurrent floor: %d\nCurrent Direction: %d\n-------------------------\n", liftLogic.GetCurrentFloor(), liftLogic.GetCurrentDirection())
			time.Sleep(time.Second) // Log every second
		}
	}()

	http.ListenAndServe(":8080", r)
}
