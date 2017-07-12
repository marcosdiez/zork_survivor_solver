package main

import "fmt"

// SurvivorPoint is a point in a SurvivorBoard
type SurvivorPoint struct {
	X int
	Y int
}

// NewSurvivorPoint creates a new SurvivorPoint
func NewSurvivorPoint(X int, Y int) (*SurvivorPoint, error) {
	if IsPointValid(X, Y) {
		return NewSurvivorPointUnsafe(X, Y), nil
	}
	theError := fmt.Errorf("error creating new Point. Invalid values: {X: %v, Y:%v}", X, Y)
	return &SurvivorPoint{X: 0, Y: 0}, theError
}

// NewSurvivorPointUnsafe creates a new SurvivorPoint, even if it is not valid
func NewSurvivorPointUnsafe(X int, Y int) *SurvivorPoint {
	return &SurvivorPoint{X: X, Y: Y}
}

// IsPointValid is boolean
func IsPointValid(X int, Y int) bool {
	return X >= 0 && X < boardSize && Y >= 0 && Y < boardSize
}

func (sp *SurvivorPoint) String() string {
	return fmt.Sprintf("[Point: X: %v, Y: %v]", sp.X, sp.Y)
}
