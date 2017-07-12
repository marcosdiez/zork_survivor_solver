package main

import "fmt"

const boardSize = 4

func navigate(sb SurvivorBoard, point *SurvivorPoint) {
	sb.step++
	sb.Board[point.X][point.Y] = sb.step
	if sb.step == ((boardSize * boardSize) - 1) {
		fmt.Println(&sb)
		return
	}
	as := sb.AvailableSteps(point)
	for _, step := range as.points {
		navigate(sb, step)
	}
}

func main() {
	fmt.Printf("Return to Zork Survivor's Solver.\n\n")

	fmt.Printf("Survivor is a game for two players, played after The Great Diffusion.\n" +
		"The players play as either Wizard Trembyle or Canuk.\n" +
		"The Wizard must move on square back or forward in one direction, and two in the other.\n" +
		"When the Wizard moves off a square it caves in, forming a pit.\n" +
		"Canuk can move to any unoccupied square on the board.\n" +
		"The object of the game for the Wizard is to make all the squares pits apart from the two occupied squares.\n" +
		"The object for Canuk is to force the Wizard to move into a pit.\n" +
		"The Wizard may,however, pass his movement turn.\n\n" +
		"Source: http://zork.wikia.com/wiki/Survivor\n\n" +

		"The numbers shows the order of squares that the Wizard should pick.\n" +
		"The Wizard starts on square 1.\n" +
		"Square 0 is where the Canuk will be in the end of the game.\n\n" +
		"Square 15 is where the Wizard will be in the end of the game.\n\n")

	fmt.Printf("Board Size: [%v]\n", boardSize)

	sb := NewSurvivorBoard()

	point, _ := NewSurvivorPoint(0, 0)
	fmt.Println(point)
	navigate(*sb, point)
}
