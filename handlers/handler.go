package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MEHUL25/lift-management-system/logic"
)

// LiftRequest represents the structure of the POST request body.
type LiftRequest struct {
	CurrentFloor     int `json:"current_floor"`
	DestinationFloor int `json:"destination_floor"`
}

// LiftResponse represents the structure of the response.
type LiftResponse struct {
	Message string `json:"message"`
}

// HandleLiftRequest handles POST requests to /api/lift.
func HandleLiftRequest(logic *logic.LiftLogic) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body
		var req LiftRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Call the lift logic to handle the request
		message := logic.HandleRequest(req.CurrentFloor, req.DestinationFloor)

		// Send a response
		resp := LiftResponse{Message: message}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
