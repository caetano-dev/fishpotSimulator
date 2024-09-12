// terminal-animation/fishpot.go
package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Fishpot struct {
	Width  int
	Height int
	Fish   []*Fish
}

func NewFishpot(width, height int, fish []*Fish) *Fishpot {
	return &Fishpot{
		Width:  width,
		Height: height,
		Fish:   fish,
	}
}

func (fp *Fishpot) FlushTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (fp *Fishpot) Draw() {
	for _, fish := range fp.Fish {
		fmt.Printf("\033[%d;%dH%s", fish.VerticalPos+1, fish.HorizontalPos+1, fish.GetFishString())
	}
	fmt.Printf("\033[%d;%dH", fp.Height+2, 0) // Move cursor to the bottom after drawing
}
