package main

import (
	"testing"
)

func TestNewBubble(t *testing.T) {
	width, height := 10, 10
	bubble := NewBubble(width, height)

	if bubble.HorizontalPos < 0 || bubble.HorizontalPos >= width {
		t.Errorf("Bubble HorizontalPos out of bounds: %d", bubble.HorizontalPos)
	}
	if bubble.VerticalPos != height-1 {
		t.Errorf("Bubble VerticalPos should be at the bottom: %d", bubble.VerticalPos)
	}
}

func TestBubbleMove(t *testing.T) {
	width, height := 10, 10
	bubble := NewBubble(width, height)

	initialVerticalPos := bubble.VerticalPos

	bubble.Move()

	if bubble.VerticalPos != initialVerticalPos-1 {
		t.Errorf("Bubble did not move up correctly")
	}
}
