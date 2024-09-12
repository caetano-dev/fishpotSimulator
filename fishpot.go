package main

import "fmt"

type Fishpot struct {
	Width  int
	Height int
	Fish   *Fish
}

func NewFishpot(width, height int, fish *Fish) *Fishpot {
	return &Fishpot{
		Width:  width,
		Height: height,
		Fish:   fish,
	}
}

func (fp *Fishpot) FlushTerminal() {
	fmt.Print("\033[H\033[2J")
}

func (fp *Fishpot) Draw() {
	for i := 0; i < fp.Height; i++ {
		for j := 0; j < fp.Width; j++ {
			if i == fp.Fish.VerticalPos && j == fp.Fish.HorizontalPos {
				fmt.Print(fp.Fish.GetFishString())
				j += len(fp.Fish.GetFishString()) - 1 // Skip the length of the fish
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
