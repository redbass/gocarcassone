package main

import "fmt"

type game struct {
	deck  deck
	board [][]*card
}

func createGame(d deck) game {
	l := len(d.cards) + len(d.rivers)
	b := make([][]*card, l)

	for i := 0; i < l; i++ {
		b[i] = make([]*card, l)
	}

	return game{
		deck:  d,
		board: b,
	}
}

func (g game) initGame() {
	g.board[0][0] = &g.deck.rivers[0]
	g.board[0][1] = &g.deck.rivers[1]
}

func (g game) draw() string {
	sBoard := ""
	for y := 0; y < len(g.board); y++ {
		for x := 0; x < len(g.board); x++ {
			p := g.board[x][y]
			if p != nil {
				sBoard += fmt.Sprintf("%s", g.board[x][y].draw())
			}
		}
		sBoard += "\n"
	}
	return sBoard
}
