package service

import "fmt"

// HandleRequest handles lift control logic.
func (l *Lift) HandleRequest(sourceFloor int, destinationFloor int) string {
	l.addFloor(sourceFloor, destinationFloor)
	return fmt.Sprintf("Request received. SourceFloor : %d. Destination floor: %d", sourceFloor, destinationFloor)
}
