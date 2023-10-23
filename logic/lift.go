package logic

import "fmt"

type LiftLogic struct {
	currentFloor int
	direction    int // 1-up and 0-down and -1 - idle
}

// NewLiftLogic creates a new instance of LiftLogic.
func NewLiftLogic() *LiftLogic {
	return &LiftLogic{
		currentFloor: 0, // Default value being ground floor
		direction:    -1,
	}
}

// HandleRequest handles lift control logic.
func (l *LiftLogic) HandleRequest(currentFloor, destinationFloor int) string {
	// Implement your lift control logic here
	// Update the lift's state, calculate the optimal course, etc.

	// For this example, we'll just acknowledge the request and update the current floor.
	l.currentFloor = destinationFloor
	return fmt.Sprintf("Request received. Current floor: %d, Destination floor: %d", currentFloor, destinationFloor)
}

// GetCurrentFloor returns the current floor of the lift.
func (l *LiftLogic) GetCurrentFloor() int {
	return l.currentFloor
}

func (l *LiftLogic) GetCurrentDirection() int {
	return l.direction
}
