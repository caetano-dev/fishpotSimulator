package fishpot

import (
	b "fishpot_simulator/bubble"
	f "fishpot_simulator/fish"
	"testing"
)

func TestNewFishpot(t *testing.T) {
	width, height := 10, 10
	fish := []*f.Fish{f.NewFish(width, height)}
	fishpot := NewFishpot(width, height, fish)

	if fishpot.Width != width {
		t.Errorf("Expected width %d, got %d", width, fishpot.Width)
	}
	if fishpot.Height != height {
		t.Errorf("Expected height %d, got %d", height, fishpot.Height)
	}
	if len(fishpot.Fish) != 1 {
		t.Errorf("Expected 1 fish, got %d", len(fishpot.Fish))
	}
}

func TestUpdateBubbles(t *testing.T) {
	width, height := 10, 10
	fishpot := NewFishpot(width, height, []*f.Fish{})

	// Add a bubble almost at the top
	fishpot.Bubbles = append(fishpot.Bubbles, &b.Bubble{HorizontalPos: 0, VerticalPos: 1})

	fishpot.UpdateBubbles()

	if len(fishpot.Bubbles) != 0 {
		t.Errorf("Expected 0 bubbles, got %d", len(fishpot.Bubbles))
	}

	// Add a bubble not at the top
	fishpot.Bubbles = append(fishpot.Bubbles, &b.Bubble{HorizontalPos: 0, VerticalPos: 2})

	fishpot.UpdateBubbles()

	if len(fishpot.Bubbles) != 1 {
		t.Errorf("Expected 1 bubble, got %d", len(fishpot.Bubbles))
	}
	if fishpot.Bubbles[0].VerticalPos != 1 {
		t.Errorf("Expected bubble to move up to 1, got %d", fishpot.Bubbles[0].VerticalPos)
	}
}
