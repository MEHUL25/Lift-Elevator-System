package main

import (
	"net/http"

	handlers "github.com/MEHUL25/lift-management-system/approach_1/handlers"
	"github.com/MEHUL25/lift-management-system/approach_1/service"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	// Approach_1

	// Create an instance of the lift logic
	l1 := service.NewApproachLift1()

	// Define routes
	r.Post("/api/reachFloor", handlers.HandleRequestFromLift(l1))
	r.Post("/api/callLift", handlers.HandleLiftCallRequest(l1))

	// Start the lift processing
	l1.StartLift()

	// Display current floor
	l1.DisplayCurrentFloor()

	http.ListenAndServe(":8000", r)
}
