package service

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Lift struct {
	currentFloor int
	direction    int // 1 - up, 0 - down, -1 - idle
	liftFloorSet *IntOrderedSet
	lock         sync.Mutex
}

// NewLift creates a new instance of Lift.
func NewApproachLift1() *Lift {
	return &Lift{
		currentFloor: 0, // Default value being ground floor
		direction:    -1,
		liftFloorSet: NewIntOrderedSet(), // Explicitly initialize the queue
	}
}

func (l *Lift) addFloor(floor int) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.liftFloorSet.Add(floor)
}

func (l *Lift) removeRequest(floor int) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.liftFloorSet.Remove(floor)
}

func (l *Lift) getNextFloor() (floor int, exists bool) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.liftFloorSet.Size() == 0 {
		return 0, false
	}

	floors := l.liftFloorSet.Items()

	fmt.Printf("data : %v\n", floors)

	switch l.direction {
	case 1: // Upwards
		idx := sort.Search(len(floors), func(i int) bool {
			return floors[i] > l.currentFloor
		})
		if idx < len(floors) {
			return floors[idx], true
		} else {
			// No floors above, check below
			if l.liftFloorSet.Size() > 0 {
				l.direction = 0
				return floors[idx], true
			}
		}
	case 0: // Downwards
		idx := sort.Search(len(floors), func(i int) bool {
			return floors[i] >= l.currentFloor
		})
		if idx > 0 {
			return floors[idx-1], true
		} else {
			// No floors below, check above
			if l.liftFloorSet.Size() > 0 {
				l.direction = 1
				return floors[idx], true
			}
		}
	case -1:
		if l.liftFloorSet.Size() > 0 {
			if floors[0] > l.currentFloor {
				l.direction = 1
			} else {
				l.direction = 0
			}
			return floors[0], true
		}
	}
	return 0, false
}

func (l *Lift) StartLift() {
	go func() {
		for {
			nextFloor, exists := l.getNextFloor()
			if !exists {
				l.direction = -1
				time.Sleep(1 * time.Second)
				continue
			}

			if nextFloor > l.currentFloor {
				l.direction = 1
			} else if nextFloor < l.currentFloor {
				l.direction = 0
			}

			if nextFloor == l.currentFloor {
				l.removeRequest(l.currentFloor)
				time.Sleep(3 * time.Second) // Wait for 2 seconds at the floor
			} else {
				l.currentFloor += l.direction*2 - 1
				l.removeRequest(l.currentFloor)
				time.Sleep(2 * time.Second)
			}
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

func (l *Lift) GetCurrentQueue() *IntOrderedSet {
	return l.liftFloorSet
}

func (l *Lift) GetCurrentDirection() int {
	return l.direction
}
