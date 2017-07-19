package main

import (
	"fmt"
	"sync"
	"time"
)

const boardSize = 4

var numUnsuccessfullAttempts int
var numSuccessfullAttempts int
var sucMux sync.Mutex
var unSucMux sync.Mutex
var waitGroup sync.WaitGroup

var runInParallel = boardSize > 5 // less than that it does not make sense

func navigate(sb SurvivorBoard, point *SurvivorPoint) {
	if runInParallel {
		defer waitGroup.Done()
	}
	sb.step++
	sb.Board[point.X][point.Y] = sb.step
	if sb.step == ((boardSize * boardSize) - 1) {
		if runInParallel {
			sucMux.Lock()
		}
		numSuccessfullAttempts++
		if runInParallel {
			sucMux.Unlock()
		}
		fmt.Println(&sb)
		return
	}
	as := sb.AvailableSteps(point)

	if as.Length() == 0 {
		if runInParallel {
			unSucMux.Lock()
		}
		numUnsuccessfullAttempts++
		if runInParallel {
			unSucMux.Unlock()
		}
	} else {
		for _, step := range as.points {
			if runInParallel {
				waitGroup.Add(1)
				go navigate(sb, step)
			} else {
				navigate(sb, step)
			}
		}
	}
}

func main() {
	start := time.Now()
	fmt.Printf("Return to Zork Survivor's Solver.\n\n")

	fmt.Printf("It's called Survivor, in honor of Wizard Trembyle and Canuk.\n" +
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
	if runInParallel {
		waitGroup.Add(1)
		go navigate(*sb, point)
		waitGroup.Wait()
	} else {
		navigate(*sb, point)
	}
	fmt.Printf("Total: %v answers and %v unsuccesful attempts\n", numSuccessfullAttempts, numUnsuccessfullAttempts)
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s\n", elapsed)
}
