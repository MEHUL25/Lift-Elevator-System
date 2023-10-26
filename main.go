package main

import (
	"net/http"

	handlers1 "github.com/MEHUL25/lift-management-system/approach_1/handlers"
	service1 "github.com/MEHUL25/lift-management-system/approach_1/service"
	handlers2 "github.com/MEHUL25/lift-management-system/approach_2/handlers"
	service2 "github.com/MEHUL25/lift-management-system/approach_2/service"
	"github.com/go-chi/chi"
)

func runApproach1(r *chi.Mux) {

	// Create an instance of the lift logic
	l := service1.NewApproachLift1()

	// Define routes
	r.Post("/api/reachFloor", handlers1.HandleRequestFromLift(l))
	r.Post("/api/callLift", handlers1.HandleLiftCallRequest(l))

	// Start the lift processing
	l.StartLift()

	// Display current floor
	l.DisplayCurrentFloor()
}

func runApproach2(r *chi.Mux) {

	// Create an instance of the lift logic
	l := service2.NewApproachLift2()

	// Define routes
	r.Post("/api/lift", handlers2.HandleRequest(l))

	// Start the lift processing
	l.StartLift()

	// Display current floor
	l.DisplayCurrentFloor()
}

func main() {
	r := chi.NewRouter()

	// Approach_1
	runApproach1(r)

	http.ListenAndServe(":8000", r)
}
