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

	fmt.Printf("It's called Survivor, in ownor of Wizard Trembyle and Canuk.\n" +
		"The only two survivors of The Great Diffusion.\n" +
		"That happend over 400 years ago.\n" +
		"You've got 16 squares.\n" +
		"I play the Wizard and you can play the Canuk.\n" +
		"So the wizard has got to move one square strait and one square diagonal in each move.\n" +
		"It can move diagonal first and strait afterwards.\n" +
		"It doesn't matter.\n" +
		"And the moment the Wizard moves off a square, it caves in and forms a pit.\n" +
		"Now Canuk can jump to any unoccupied square and far sure to move into a pit and he's won.\n" +
		"And the object is to have all the squares turned into pits, except the one you will be sitting on.\n" +
		"That takes strategy.\n\n" +

		"The numbers shows the order of squares that the Wizard should pick.\n" +
		"The Wizard starts on square 1.\n" +
		"Square 15 is where the Wizard will be in the end of the game.\n" +
		"Square 0 is where the Canuk will be in the end of the game.\n\n")

	fmt.Printf("Board Size: [%v]\n", boardSize)

	sb := NewSurvivorBoard()

	point, _ := NewSurvivorPoint(0, 0)
	fmt.Println(point)
	navigate(*sb, point)
}
