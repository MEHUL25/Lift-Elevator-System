package main

import (
	"fmt"
	"net/http"

	handlers1 "github.com/MEHUL25/lift-management-system/approach_1/handlers"
	service1 "github.com/MEHUL25/lift-management-system/approach_1/service"
	handlers2 "github.com/MEHUL25/lift-management-system/approach_2/handlers"
	service2 "github.com/MEHUL25/lift-management-system/approach_2/service"
	handlers3 "github.com/MEHUL25/lift-management-system/multipleLifts/handlers"
	service3 "github.com/MEHUL25/lift-management-system/multipleLifts/service"
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

func runMultipleLifts(r *chi.Mux) {
	liftManager := service3.NewLiftManager(4)

	r.Post("/api/callLift", handlers3.HandleLiftCallRequest(liftManager))

	for _, lift := range liftManager.Lifts {
		lift.StartLift()
		lift.DisplayCurrentFloor()
	}

	for _, lift := range liftManager.Lifts {
		endpoint := fmt.Sprintf("/api/lift/%d/request", lift.LiftID)
		r.Post(endpoint, handlers3.HandleRequestFromLift(lift.LiftID, lift))
		lift.StartLift()
		lift.DisplayCurrentFloor()
	}
}

func main() {
	r := chi.NewRouter()

	// Approach_1
	runApproach1(r)

	http.ListenAndServe(":8000", r)
}
