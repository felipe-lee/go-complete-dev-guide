package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	expectedLength := 52
	if len(d) != expectedLength {
		t.Errorf("Expected deck length of %d, but got %d", expectedLength, len(d))
	}

	expectedFirstCard := "Ace of Spades"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card to be %v, but got %v", expectedFirstCard, d[0])
	}

	expectedLastCard := "King of Clubs"
	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last card to be %v, but got %v", expectedLastCard, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_deck_test_file.txt"

	if _, err := os.Stat(filename); err == nil {
		err = os.Remove(filename)

		if err != nil {
			t.Errorf("Error removing file: %v", err)
		}
	}

	d := newDeck()
	err := d.saveToFile(filename)

	if err != nil {
		t.Errorf("Error saving deck to file: %v", err)
	}

	loadedDeck := newDeckFromFile(filename)

	expectedDeckLength := 52
	if len(loadedDeck) != expectedDeckLength {
		t.Errorf("Expected %d cards in the deck, got %d", expectedDeckLength, len(loadedDeck))
	}

	err = os.Remove(filename)

	if err != nil {
		t.Errorf("Error removing file: %v", err)
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()

	handSize := 5

	hand, remainingDeck := deal(d, handSize)

	if len(hand) != handSize {
		t.Errorf("Expected hand to be size %d, but got %d", handSize, len(hand))
	}

	expectedRemainderSize := 52 - handSize
	if len(remainingDeck) != expectedRemainderSize {
		t.Errorf("Expected to have %d cards left, but got %d", expectedRemainderSize, len(remainingDeck))
	}
}
