package main

import (
	"fmt"
)

func main() {
	d := readJSONDeck("deck.json")
	g := createGame(d)

	g.initGame()

	res := g.draw()
	fmt.Println(res)
}
