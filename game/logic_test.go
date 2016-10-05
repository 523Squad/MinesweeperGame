package game

import (
	"fmt"
	"testing"
)

func TestPointToString(t *testing.T) {
	p := point{true, true, 4, false}
	s := " true neighbours 4"
	if p.toString() != s {
		error := fmt.Sprintf("Expected %s, got %s", s, p.toString())
		t.Error(error)
	}
}
