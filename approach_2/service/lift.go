package service

import (
	"fmt"
	"sync"
	"time"
)

type Lift struct {
	currentFloor int
	direction    int // 1 - up, 0 - down, -1 - idle
	liftFloorSet *OrderedSet
	lock         sync.Mutex
}

// NewLift creates a new instance of Lift.
func NewApproachLift2() *Lift {
	return &Lift{
		currentFloor: 0, // Default value being ground floor
		direction:    -1,
		liftFloorSet: NewOrderedSet(), // Explicitly initialize the queue
	}
}

func (l *Lift) addFloor(sourceFloor, destinationFloor int) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.liftFloorSet.Add(FloorRequest{Source: sourceFloor, Destination: destinationFloor})

}

func (l *Lift) getNextRequest() (FloorRequest, bool) {
	l.lock.Lock()
	defer l.lock.Unlock()

	items := l.liftFloorSet.Items()

	// If lift is idle or there are no requests, return false
	if l.direction == -1 || len(items) == 0 {
		return FloorRequest{}, false
	}

	// If going upwards
	if l.direction == 1 {
		for _, req := range items {
			if req.Source >= l.currentFloor {
				return req, true
			}
		}
	}

	// If going downwards
	if l.direction == 0 {
		for i := len(items) - 1; i >= 0; i-- {
			if items[i].Source <= l.currentFloor {
				return items[i], true
			}
		}
	}

	return FloorRequest{}, false
}

func (l *Lift) StartLift() {
	go func() {
		for {
			req, exists := l.getNextRequest()

			if !exists {
				l.direction = -1
				time.Sleep(1 * time.Second)
				continue
			}

			// Update direction based on destination
			if req.Destination > l.currentFloor {
				l.direction = 1
			} else {
				l.direction = 0
			}

			// Simulate lift movement
			if req.Source != l.currentFloor {
				time.Sleep(2 * time.Second) // Assume 2 seconds to reach each floor
				l.currentFloor = req.Source
			}

			if req.Destination != l.currentFloor {
				time.Sleep(2 * time.Second)
				l.currentFloor = req.Destination
			}

			// Remove served request
			l.liftFloorSet.Remove(req)
		}
	}()
}

func (l *Lift) DisplayCurrentFloor() {
	go func() {
		for {
			fmt.Printf("\n-------------------------\nCurrent floor: %d\nCurrent Direction: %d\nCurrent Queue: %v\n-------------------------\n", l.GetCurrentFloor(), l.GetCurrentDirection(), l.GetCurrentQueue())
			time.Sleep(1 * time.Second)
		}
	}()
}

// GetCurrentFloor returns the current floor of the lift.
func (l *Lift) GetCurrentFloor() int {
	return l.currentFloor
}

func (l *Lift) GetCurrentQueue() *OrderedSet {
	return l.liftFloorSet
}

func (l *Lift) GetCurrentDirection() int {
	return l.direction
}
