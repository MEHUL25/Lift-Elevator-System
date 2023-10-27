package service

import "math"

type LiftManager struct {
	Lifts []*Lift
}

func NewLiftManager(numLifts int) *LiftManager {
	lm := &LiftManager{
		Lifts: make([]*Lift, numLifts),
	}
	for i := 0; i < numLifts; i++ {
		lm.Lifts[i] = NewApproachLift1(i + 1) // LiftID starts from 1
	}
	return lm
}

func (lm *LiftManager) AssignLift(requestFloor int) *Lift {
	var selectedLift *Lift
	minDistance := int(^uint(0) >> 1) // Max int value

	for _, lift := range lm.Lifts {
		dif := float64((lift.GetCurrentFloor()) - (requestFloor))
		distance := int(math.Abs(dif))
		if distance < minDistance {
			selectedLift = lift
			minDistance = distance
		}
	}
	return selectedLift
}
