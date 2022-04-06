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
	if c.Value == 1 {
		return fmt.Sprintf("A %s", c.Seed)
	}

	if c.Value < 11 {
		return fmt.Sprintf("%d %s", c.Value, c.Seed)
	}

	semi := map[int]string{1: "A", 11: "J", 12: "Q", 13: "K"}
	return fmt.Sprintf("%s %s", semi[c.Value], c.Seed)
}

func CreateDeck() [52]*Card {
	seeds := []string{"Cuori", "Quadri", "Picche", "Fiori"}
	var deck [52]*Card
	i := 0
	for _, c := range seeds {
		for num := 1; num <= 13; num++ {
			deck[i] = New(num, c)
			i++
		}
	}
	//Mescola
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return deck
}

func main() {
	var deck = CreateDeck()
	for _, value := range deck {
		fmt.Println(value.String())
	}
}
