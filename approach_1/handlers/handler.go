package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MEHUL25/lift-management-system/approach_1/service"
)

// LiftRequest represents the structure of the POST request body.
type LiftDestinationRequest struct {
	DestinationFloor int `json:"destination_floor"`
}

type LiftCallingRequest struct {
	CurrentFloor int `json:"current_floor"`
}

// LiftResponse represents the structure of the response.
type LiftResponse struct {
	Message string `json:"message"`
}

// HandleLiftRequest handles POST requests to /api/lift.
func HandleRequestFromLift(logic *service.Lift) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body
		var req LiftDestinationRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Hello\n\n")

		// Call the lift logic to handle the request
		message := logic.HandleRequestFromLift(req.DestinationFloor)

		// Send a response
		resp := LiftResponse{Message: message}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

// HandleLiftRequest handles POST requests to /api/lift.
func HandleLiftCallRequest(logic *service.Lift) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body
		var req LiftCallingRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Hello\n\n")

		// Call the lift logic to handle the request
		message := logic.HandleLiftCallRequest(req.CurrentFloor)

		// Send a response
		resp := LiftResponse{Message: message}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
