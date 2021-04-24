package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck size of 52, but got %d", len(d))
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	size := 5

	hand, _ := deal(d, size)

	if len(hand) != size {
		t.Errorf("Expected a deal of %d cards", len(hand))
	}
}
