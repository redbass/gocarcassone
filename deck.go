package main

type deck struct {
	cards []card
	rivers []card
}

func newDeck(cards [][]string, rivers [][]string) deck {
	var deckCards []card
	var deckRivers []card

	for _, c := range cards {
		deckCards = append(deckCards, makeCard(c))
	}
	for _, r := range rivers {
		deckRivers = append(deckRivers, makeCard(r))
	}
	return deck{
		cards: deckCards,
		rivers: deckRivers,
	}
}
