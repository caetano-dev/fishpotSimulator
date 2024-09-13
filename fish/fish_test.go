package fish

import (
	"testing"
)

func TestNewFish(t *testing.T) {
	width, height := 10, 10
	fish := NewFish(width, height)

	if fish.HorizontalPos < 0 || fish.HorizontalPos >= width {
		t.Errorf("Fish HorizontalPos out of bounds: %d", fish.HorizontalPos)
	}
	if fish.VerticalPos < 0 || fish.VerticalPos >= height {
		t.Errorf("Fish VerticalPos out of bounds: %d", fish.VerticalPos)
	}
	if fish.Direction != "right" && fish.Direction != "left" {
		t.Errorf("Fish Direction invalid: %s", fish.Direction)
	}
}

func TestFishMove(t *testing.T) {
	width, height := 10, 10
	fish := NewFish(width, height)

	initialHorizontalPos := fish.HorizontalPos
	initialVerticalPos := fish.VerticalPos

	fish.Move()
	fish.Move()

	if fish.HorizontalPos == initialHorizontalPos && fish.VerticalPos == initialVerticalPos {
		t.Errorf("Fish did not move")
	}
}

func TestFishChangeDirection(t *testing.T) {
	fish := &Fish{Direction: "right"}
	fish.changeDirection()
	if fish.Direction != "left" {
		t.Errorf("Expected direction to be left, got %s", fish.Direction)
	}

	fish.changeDirection()
	if fish.Direction != "right" {
		t.Errorf("Expected direction to be right, got %s", fish.Direction)
	}
}

func TestGetFishString(t *testing.T) {
	fish := &Fish{Direction: "right"}
	if fish.GetFishString() != rightFish {
		t.Errorf("Expected %s, got %s", rightFish, fish.GetFishString())
	}

	fish.Direction = "left"
	if fish.GetFishString() != leftFish {
		t.Errorf("Expected %s, got %s", leftFish, fish.GetFishString())
	}
}
