package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type game []string

func printState() string {
	return "California"
}

func newGame() game {
	players := []string{"Zero", "Cross"}
	gridValue := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	moves := game{}

	for _, suit := range players {
		for _, value := range gridValue {
			moves = append(moves, value+" of "+suit)
		}
	}

	return moves
}

func deal(d game, handSize int) (game, game) {
	return d[:handSize], d[handSize:]
}

func (d game) toString() string {
	return strings.Join([]string(d), ",")
}

func (d game) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newGameFromFile(filename string) game {
	ds, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	ss := strings.Split(string(ds), ",")
	return game(ss)
}

func (d game) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn((len(d) - 1))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d game) print() {
	fmt.Println(d)

	for i, card := range d {
		println(i, card)
	}
}
