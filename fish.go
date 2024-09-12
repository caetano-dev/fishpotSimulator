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
		HorizontalPos: 0,
		VerticalPos:   height / 2,
		Direction:     "right",
		Width:         width,
		Height:        height,
	}
}

func (f *Fish) Move() {
	// Randomly move up or down
	if rand.Intn(2) == 0 {
		if rand.Intn(2) == 0 && f.VerticalPos > 0 {
			f.VerticalPos--
		} else if f.VerticalPos < f.Height-1 {
			f.VerticalPos++
		}
	}

	// Randomly move left or right
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

func (f *Fish) GetFishString() string {
	if f.Direction == "right" {
		return rightFish
	}
	return leftFish
}
