package costructor

import (
	"math/rand"
	"time"
)

type Card struct {
	Value int
	Seed  string
}

type Deck struct {
	Cards []*Card
}

//Costruttore per struttura Card
func New(v int, s string) *Card {
	return &Card{Value: v, Seed: s}
}

//Costruttore struttura Deck
func CreateDeck(start int) *Deck {
	seeds := []string{"♥", "♦", "♠", "♣"}
	var deck []*Card
	for _, c := range seeds {
		for num := start; num <= 14; num++ {
			deck = append(deck, New(num, c))
		}
	}
	//Mescola il mazzo di carte
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return &Deck{Cards: deck}
}

func (d *Deck) DrawHand(n int) []*Card {
	temp := d.Cards[:n]
	d.Cards = d.Cards[n:]
	return temp
}
