package main

import (
	"math/rand"
)

type Bubble struct {
	HorizontalPos int
	VerticalPos   int
}

func NewBubble(width, height int) *Bubble {
	return &Bubble{
		HorizontalPos: rand.Intn(width),
		VerticalPos:   height - 1,
	}
}

func (b *Bubble) Move() {
	if b.VerticalPos > 0 {
		b.VerticalPos--
	}
}
