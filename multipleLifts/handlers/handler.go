package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MEHUL25/lift-management-system/multipleLifts/service"
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
func HandleRequestFromLift(liftID int, logic *service.Lift) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LiftDestinationRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if logic.LiftID != liftID {
			http.Error(w, "Wrong lift endpoint accessed.", http.StatusBadRequest)
			return
		}

		message := logic.HandleRequestFromLift(req.DestinationFloor)

		resp := LiftResponse{Message: message}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleLiftCallRequest(manager *service.LiftManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LiftCallingRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		assignedLift := manager.AssignLift(req.CurrentFloor)
		assignedLift.AddFloor(req.CurrentFloor)

		resp := LiftResponse{Message: fmt.Sprintf("Request received. Lift %d assigned.", assignedLift.LiftID)}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}
