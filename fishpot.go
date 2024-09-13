package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
)

type Fishpot struct {
	Width   int
	Height  int
	Fish    []*Fish
	Bubbles []*Bubble
}

func NewFishpot(width, height int, fish []*Fish) *Fishpot {
	return &Fishpot{
		Width:   width,
		Height:  height,
		Fish:    fish,
		Bubbles: []*Bubble{},
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
	for _, bubble := range fp.Bubbles {
		fmt.Printf("\033[%d;%dH%s", bubble.VerticalPos+1, bubble.HorizontalPos+1, "o")
	}
}

func (fp *Fishpot) Update() {
	for _, fish := range fp.Fish {
		fish.Move()
	}
	fp.UpdateBubbles()
}

func (fp *Fishpot) UpdateBubbles() {
	// Move existing bubbles
	for _, bubble := range fp.Bubbles {
		bubble.Move()
	}

	// Remove bubbles that have reached the top
	fp.Bubbles = filterBubbles(fp.Bubbles, func(bubble *Bubble) bool {
		return bubble.VerticalPos == 0
	})

	// Randomly create new bubbles
	if rand.Intn(10) < 2 {
		fp.Bubbles = append(fp.Bubbles, NewBubble(fp.Width, fp.Height))
	}
}

func filterBubbles(bubbles []*Bubble, hasReachedTop func(*Bubble) bool) (remaining []*Bubble) {
	for _, bubble := range bubbles {
		if !hasReachedTop(bubble) {
			remaining = append(remaining, bubble)
		}
	}
	return
}
