package deal

import (
	"strconv"
	"testing"
)

func TestColorMatch(t *testing.T) {
	a := uint(0)
	b := uint(0+FIGURES)
	result := colorMatch(a, b)
	expected := false
	if result != expected {
		t.Errorf("TestColorMatch was incorrect, got: %s, want: %s.",
			strconv.FormatBool(result), strconv.FormatBool(expected))
	}
	a = uint(FIGURES-1)
	b = uint(FIGURES)
	result = colorMatch(a, b)
	expected = false
	if result != expected {
		t.Errorf("TestColorMatch was incorrect, got: %s, want: %s.",
			strconv.FormatBool(result), strconv.FormatBool(expected))
	}
	a = uint(0)
	b = uint(FIGURES-1)
	result = colorMatch(a, b)
	expected = true
	if result != expected {
		t.Errorf("TestColorMatch was incorrect, got: %s, want: %s.",
			strconv.FormatBool(result), strconv.FormatBool(expected))
	}
}
