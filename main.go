package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Value int
	Seed   string
}

func New(v int, s string) *Card {
	return &Card{Value: v, Seed: s}
}

func (c Card) String() string {

	if c.Value > 1 && c.Value < 11 {
		return fmt.Sprintf("%3d di %5s", c.Value, c.Seed)
	}

	seeds := map[int]string{1: "A", 11: "J", 12: "Q", 13: "K"}
	return fmt.Sprintf("%2s di %5s", seeds[c.Value], c.Seed)
}

func CreateDeck() []*Card {
	seeds := []string{"Cuori", "Quadri", "Picche", "Fiori"}
	var deck []*Card
	for _, c := range seeds {
		for num := 1; num <= 13; num++ {
			deck = append(deck, New(num, c))
		}
	}
	//Mescola il mazzo di carte
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

func DrawHand(deck []*Card) ([]*Card, []*Card) {
	temp := deck[:5]
	deck = deck[5:]
	return temp , deck
}

func Print(hand []*Card) {
	for _, h:= range hand {
		fmt.Print(h.String())
	}
}

func main() {

	var deck = CreateDeck()

	hand, deck := DrawHand(deck)
	
	Print(hand)
//	for _, value := range deck {
//		fmt.Println(value.String())
//	}
}
