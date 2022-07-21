package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 56 {
		t.Errorf("Expected 56 but recieved %v", len(d))
	}
}
