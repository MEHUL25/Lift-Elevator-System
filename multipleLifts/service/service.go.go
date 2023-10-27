package service

import "fmt"

// HandleRequest handles lift control logic.
func (l *Lift) HandleRequestFromLift(destinationFloor int) string {
	l.AddFloor(destinationFloor)
	return fmt.Sprintf("Request received.  Destination floor: %d", destinationFloor)
}

// HandleRequest handles lift control logic.
func (l *Lift) HandleLiftCallRequest(sourceFloor int) string {
	l.AddFloor(sourceFloor)
	return fmt.Sprintf("Request received. Current floor: %d : ", sourceFloor)
}
