package main

import (
	"os"
	"testing"
)

func TestNewGame(t *testing.T) {
	d := newGame()

	if len(d) != 16 {
		t.Errorf("Expected game length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Club" {
		t.Errorf("Expected first card 'Four of Club', but got %v", d[0])
	}
}

func TestSaveToGameAndNewGameFromFile(t *testing.T) {
	game := newGame()
	game.saveToFile("_gametesting")

	loadedGame := newGameFromFile("_gametesting")

	if len(loadedGame) != 16 {
		t.Errorf("Expected game length of 16, but got %v", len(loadedGame))
	}

	os.Remove("_gametesting")
}
