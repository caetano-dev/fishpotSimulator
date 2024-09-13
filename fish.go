package main

import (
	"math/rand"
	"time"
)

const (
	rightFish = "><(((('>"
	leftFish  = "<'(((><"
)

type Fish struct {
	HorizontalPos int
	VerticalPos   int
	Direction     string
	Width         int
	Height        int
}

func NewFish(width, height int) *Fish {
	rand.NewSource(time.Now().UnixNano())
	return &Fish{
		HorizontalPos: rand.Intn(width),
		VerticalPos:   rand.Intn(height),
		Direction:     "right",
		Width:         width,
		Height:        height,
	}
}

func (f *Fish) Move() {
	f.moveVertically()
	f.moveHorizontally()
}

func (f *Fish) moveVertically() {
	// Randomly move up or down
	if rand.Intn(2) == 0 {
		if rand.Intn(2) == 0 && f.VerticalPos > 0 {
			f.VerticalPos--
		} else if f.VerticalPos < f.Height-1 {
			f.VerticalPos++
		}
	}
}

func (f *Fish) moveHorizontally() {
	if rand.Intn(50) == 0 {
		f.changeDirection()
	}

	//check if fish is at the edge
	if f.Direction == "right" {
		if f.HorizontalPos < f.Width-len(rightFish) {
			f.HorizontalPos++
		} else {
			f.Direction = "left"
		}
	} else {
		if f.HorizontalPos > 0 {
			f.HorizontalPos--
		} else {
			f.Direction = "right"
		}
	}
}

func (f *Fish) changeDirection() {
	f.Direction = map[string]string{"right": "left", "left": "right"}[f.Direction]
}

func (f *Fish) GetFishString() string {
	return map[string]string{"right": rightFish, "left": leftFish}[f.Direction]
}
