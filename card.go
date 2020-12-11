package main

import (
	"strings"
)

const borderLen = 3
const dataLen = borderLen * borderLen

var cardinalIndex = map[string]int{
	"N": 1,
	"W": 5,
	"S": 7,
	"E": 3,
	"C": 4,
}
var borderTypes = map[string]string{
	"city":      "C",
	"road":      "S",
	"river":     "R",
	"farm":      "F",
	"monastery": "M",
}

type card struct {
	data       []string
	rotation   int
}

func makeCard(d []string) card {
	return card{data: d, rotation: 0}
}

func (c card) getData(d int) string {
	return c.data[d]
}

func (c card) draw() string {
	scale := 2

	sCard := ""
	sRow := ""

	typeColor := map[string]string{
		"C": "\033[33m",
		"S": "\033[37m",
		"R": "\033[34m",
		"F": "\033[32m",
		"M": "\033[35m",
	}
	char := "\u258A"

	for d := 0; d < dataLen; d++ {
		color := typeColor[c.getData(d)]
		sRow += strings.Repeat(color+char, scale)
		if (d+1)%borderLen == 0 {
			sCard += strings.Repeat(sRow+"\n", scale)
			sRow = ""
		}
	}

	return sCard
}
