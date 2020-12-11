package main

const borderLen = 3

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

