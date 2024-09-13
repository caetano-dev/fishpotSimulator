package main

import (
	"fmt"
	"os"
	"time"

	f "fishpot_simulator/fish"
	fp "fishpot_simulator/fishpot"

	"golang.org/x/term"
)

func main() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		return
	}

	// Adjust width and height to leave some space for borders
	width -= 2
	height -= 2

	// Create multiple fish
	numFish := 5
	fish := make([]*f.Fish, numFish)
	for i := 0; i < numFish; i++ {
		fish[i] = f.NewFish(width, height)
	}

	fishpot := fp.NewFishpot(width, height, fish)

	for {
		fishpot.FlushTerminal()
		fishpot.Draw()
		time.Sleep(300 * time.Millisecond)
		for _, fish := range fish {
			fish.Move()
		}
		fishpot.Update()
	}
}
