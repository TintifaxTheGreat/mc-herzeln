package main

import (
	"testing"
)

func TestNewBitmap(t *testing.T) {
	b := newBitmap(true)
	expected := bitmap(4294967295)
	if *b != expected {
		t.Errorf("TestToString was incorrect, got: %s, want: %s.", b.toString(), expected.toString())
	}
}

func TestToString(t *testing.T) {
	// create helpers
	helper := new(Helper)
	CARDSTRINGS = helper.Cardstrings()

	b := bitmap(1998)
	result := b.toString()
	expected := "HK HO HU H8 H7 SA SK SO "
	if result != expected {
		t.Errorf("TestToString was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestSet(t *testing.T) {
	b := bitmap(1)
	b.set(1)
	b.set(3)
	var expected bitmap = 11
	if b != expected {
		t.Errorf("TestSet was incorrect, got: %d, want: %d.", b, expected)
	}
}

func TestUnset(t *testing.T) {
	b := bitmap(15)
	b.unset(2)
	b.unset(0)
	var expected bitmap = 10
	if b != expected {
		t.Errorf("TestUnset was incorrect, got: %d, want: %d.", b, expected)
	}
	b = bitmap(25)
	b.unset(1)
	b.unset(3)
	expected = 17
	if b != expected {
		t.Errorf("TestUnset was incorrect, got: %d, want: %d.", b, expected)
	}
}

func TestNext(t *testing.T) {
	b := bitmap(130)
	result := b.next(1)
	expected := uint(7)
	if result != expected {
		t.Errorf("TestNext was incorrect, got: %d, want: %d.", result, expected)
	}
}
