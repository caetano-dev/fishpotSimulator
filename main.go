// create a command line program that will simulate a fishpot. The fish will be represented by the character `~` and the fishpot will be represented by the character `#`. The fish will move from left to right and when it reaches the end of the fishpot, it will start again from the beginning. The fishpot will have a width of 10 characters and a height of 10 characters. The fish will move every 500 milliseconds. The fish will move 10 times before the program ends.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width     = 20
	height    = 10
	rightFish = "><(((('>"
	leftFish  = "<'(((><"
)

func flushTerminal() {
	fmt.Print("\033[H\033[2J")
}

func drawFishpot(horizontalFishPos, verticalFishPos int, direction string) {
	fmt.Println("===========================")
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == verticalFishPos && j == horizontalFishPos {
				if direction == "right" {
					fmt.Print(rightFish)
				} else {
					fmt.Print(leftFish)
				}
				j += len(rightFish) - 1 // Skip the length of the fish
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println("===========================")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	horizontalFishPos := 0
	verticalFishPos := height / 2
	direction := "right"

	for {
		flushTerminal()
		drawFishpot(horizontalFishPos, verticalFishPos, direction)
		time.Sleep(500 * time.Millisecond)

		// Randomly move up or down
		if rand.Intn(2) == 0 {
			if rand.Intn(2) == 0 && verticalFishPos > 0 {
				verticalFishPos--
			} else if verticalFishPos < height-1 {
				verticalFishPos++
			}
		}

		// Move horizontally
		if direction == "right" {
			horizontalFishPos++
			if horizontalFishPos >= width-len(rightFish) {
				horizontalFishPos = width - len(rightFish)
				direction = "left"
			}
		} else {
			horizontalFishPos--
			if horizontalFishPos < 0 {
				horizontalFishPos = 0
				direction = "right"
			}
		}
	}
}
