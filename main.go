package main

func main() {
	game := newGameFromFile("my_game")
	game.shuffle()
	game.print()

	// game.saveToFile("my_game")

	/* hand, remaining := turns(game, 4)
	hand.print()
	remaining.print() */

}
