package main

import (
	"testing"
)

func TestNewBitmap(t *testing.T) {
	b := NewBitmap(true)
	expected := Bitmap(4294967295)
	if *b != expected {
		t.Errorf("TestToString was incorrect, got: %s, want: %s.", b.ToString(), expected.ToString())
	}
}

func TestToString(t *testing.T) {
	// create helpers
	helper := new(Helper)
	CARDSTRINGS = helper.Cardstrings()

	b := Bitmap(1998)
	result := b.ToString()
	expected := "HK HO HU H8 H7 SA SK SO "
	if result != expected {
		t.Errorf("TestToString was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestSet(t *testing.T) {
	b := Bitmap(1)
	b.Set(1)
	b.Set(3)
	var expected Bitmap = 11
	if b != expected {
		t.Errorf("TestSet was incorrect, got: %d, want: %d.", b, expected)
	}
}

func TestUnset(t *testing.T) {
	b := Bitmap(15)
	b.Unset(2)
	b.Unset(0)
	var expected Bitmap = 10
	if b != expected {
		t.Errorf("TestUnset was incorrect, got: %d, want: %d.", b, expected)
	}
}

func TestNext(t *testing.T) {
	b := Bitmap(130)
	result := b.Next(1)
	expected := uint(7)
	if result != expected {
		t.Errorf("TestNext was incorrect, got: %d, want: %d.", result, expected)
	}
}
