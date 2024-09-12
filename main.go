// create a command line program that will simulate a fishpot. The fish will be represented by the character `~` and the fishpot will be represented by the character `#`. The fish will move from left to right and when it reaches the end of the fishpot, it will start again from the beginning. The fishpot will have a width of 10 characters and a height of 10 characters. The fish will move every 500 milliseconds. The fish will move 10 times before the program ends.
package main

import (
	"fmt"
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

func drawFishpot(horizontalFishPos int, direction string) {
	fmt.Println("===========================")
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if i == height/2 && j == horizontalFishPos {
				if direction == "right" {
					fmt.Print(rightFish)
				} else {
					fmt.Print(leftFish)
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println("===========================")
}

func main() {
	horizontalFishPos := 0
	direction := "right"

	for {
		flushTerminal()
		drawFishpot(

			horizontalFishPos, direction)
		time.Sleep(500 * time.Millisecond)

		if direction == "right" {
			horizontalFishPos++
			if horizontalFishPos >= width {
				horizontalFishPos = width - 1
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
