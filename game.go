package main

import (
	"strings"
)

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
	pos := [][]int{
		{0, 0}, {0, 1}, {1, 1}, {2, 1}, {3, 1}, {3, 2}, {3, 3}, {4, 3}, {5, 3}, {6, 3}, {7, 3}, {7, 4},
	}

	for i := 0; i < len(g.deck.rivers); i++ {
		p := pos[i]
		point := &g.deck.rivers[i]
		g.board[p[1]][p[0]] = point
	}
}

func (g game) draw() string {
	scale := 3
	boardLength := len(g.board)

	sBoard := ""
	for x := 0; x < boardLength; x++ {
		sBoard += drawRow(g.board[x], scale)
	}
	return sBoard
}

func drawRow(row []*card, scale int) string {
	sRow := ""
	t := ""
	m := ""
	b := ""
	for _, c := range row {
		st, sm, sb := drawCard(c, scale)
		t += st
		m += sm
		b += sb
	}

	if strings.TrimSpace(t) == ""{
		return ""
	}

	sRow += strings.Repeat(t+"\n", scale)
	sRow += strings.Repeat(m+"\n", scale)
	sRow += strings.Repeat(b+"\n", scale)
	return sRow
}

func drawCard(c *card, hScale int) (string, string, string) {
	dataLen := borderLen * borderLen

	sCard := make([]string, 3)
	sRow := ""

	typeColor := map[string]string{
		"C": "\033[33m",
		"S": "\033[37m",
		"R": "\033[34m",
		"F": "\033[32m",
		"M": "\033[35m",
	}
	for d := 0; d < dataLen; d++ {
		char := " "
		if c != nil {
			color := typeColor[c.getData(d)]
			char = color + "\u258A"
		}

		sRow += strings.Repeat(char, hScale)

		if (d+1)%borderLen == 0 {
			sCard[d/3] = sRow
			sRow = ""
		}
	}

	return sCard[0], sCard[1], sCard[2]
}
