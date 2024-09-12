// create a command line program that will simulate a fishpot. The fish can move randonly up or down and left or right.

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/term"
)

const (
	rightFish = "><(((('>"
	leftFish  = "<'(((><"
	bubble    = "0"
)

func flushTerminal() {
	fmt.Print("\033[H\033[2J")
}

func switchDirection(direction string) string {
	return map[string]string{"right": "left", "left": "right"}[direction]

}

func drawFishpot(horizontalFishPos, verticalFishPos, width, height int, direction string) {
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
}

func main() {
	rand.NewSource(time.Now().UnixNano())
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	// Adjust width and height to leave some space for borders
	width -= 2
	height -= 2

	horizontalFishPos := 0
	verticalFishPos := height / 2
	direction := "right"

	for {
		flushTerminal()
		drawFishpot(horizontalFishPos, verticalFishPos, width, height, direction)
		time.Sleep(200 * time.Millisecond)

		// Randomly move up or down
		if rand.Intn(2) == 0 {
			if rand.Intn(2) == 0 && verticalFishPos > 0 {
				verticalFishPos--
			} else if verticalFishPos < height-1 {
				verticalFishPos++
			}
		}

		// Randomly move left or right
		switchDirection(direction)
		if direction == "right" {
			if horizontalFishPos < width-len(rightFish) {
				horizontalFishPos++
			} else {
				direction = "left"
			}
		} else {
			if horizontalFishPos > 0 {
				horizontalFishPos--
			} else {
				direction = "right"
			}
		}

	}

}
