package main

func main() {
	cards := newGameFromFile("my_cards")
	cards.shuffle()
	cards.print()

	// cards.saveToFile("my_cards")

	/* hand, remaining := deal(cards, 4)
	hand.print()
	remaining.print() */
}
