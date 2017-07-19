package main

import (
	"fmt"
)

// SurvivorBoard is a a standard Return To Zork's Survival board, with 4x4 squares
type SurvivorBoard struct {
	Board [boardSize][boardSize]int
	step  int
}

// NewSurvivorBoard creates a new SurvivorBoard
func NewSurvivorBoard() *SurvivorBoard {
	return &SurvivorBoard{}
}

func (sb *SurvivorBoard) addPointIfValid(result *SurvivorPointList, X int, Y int) {
	point, err := NewSurvivorPoint(X, Y)
	if err == nil {
		if sb.Board[X][Y] == 0 {
			result.AddPoint(point)
		}
	}
}

// AvailableSteps returns the number of steps available in a point
func (sb *SurvivorBoard) AvailableSteps(sp *SurvivorPoint) *SurvivorPointList {
	result := NewSurvivorPointList()

	sb.addPointIfValid(result, sp.X+1, sp.Y+2)
	sb.addPointIfValid(result, sp.X+2, sp.Y+1)

	sb.addPointIfValid(result, sp.X-1, sp.Y+2)
	sb.addPointIfValid(result, sp.X-2, sp.Y+1)

	sb.addPointIfValid(result, sp.X+1, sp.Y-2)
	sb.addPointIfValid(result, sp.X+2, sp.Y-1)

	sb.addPointIfValid(result, sp.X-1, sp.Y-2)
	sb.addPointIfValid(result, sp.X-2, sp.Y-1)

	return result
}

func (sb *SurvivorBoard) String() string {
	result := ""
	for row := 0; row < len(sb.Board); row++ {
		result += "["
		for col := 0; col < len(sb.Board[col]); col++ {
			result += fmt.Sprintf(" %2d", sb.Board[row][col])
		}
		result += "]\n"
	}
	return result
}
