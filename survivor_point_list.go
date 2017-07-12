package main

import (
	"fmt"
)

// SurvivorPointList is a list of SurvivorPoint
type SurvivorPointList struct {
	points []*SurvivorPoint
}

// NewSurvivorPointList instantiate a new SuriviorPointList
func NewSurvivorPointList() *SurvivorPointList {
	return &SurvivorPointList{}
}

// PointExists checks if a point is in a list
func (slp *SurvivorPointList) PointExists(sp *SurvivorPoint) bool {
	for _, pointInList := range slp.points {
		if pointInList.X == sp.X && pointInList.Y == sp.Y {
			return true
		}
	}
	return false
}

// AddPoint checks if a point is in a list
func (slp *SurvivorPointList) AddPoint(sp *SurvivorPoint) error {
	if slp.PointExists(sp) {
		return fmt.Errorf("error: the point already %v exists in the PointList", sp)
	}
	slp.points = append(slp.points, sp)
	return nil
}

// String ...
func (slp SurvivorPointList) String() string {
	return fmt.Sprint(slp.points)
}

// Length return the size of the list
func (slp *SurvivorPointList) Length() int {
	return len(slp.points)
}
