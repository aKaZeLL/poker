package main

import (
	"fmt"
)

func String(n int) string {
	if n > 1 && n < 11 {
		return fmt.Sprintf("%d", n)
	}
	seeds := map[int]string{11: "J", 12: "Q", 13: "K", 14: "A"}
	return seeds[n]
}

func Print(hand []*Card) {
	for _, h := range hand {
		fmt.Printf("%s%s ", String(h.Value), h.Seed)
	}
	fmt.Println()
}

func ChangeCards(hand []*Card, deck *Deck) []*Card {
	var n int
	var c1, c2, c3, c4 int
	var indici []int

	fmt.Print("Quante carte vuoi cambiare: ")
	fmt.Scanln(&n)

	if n > 0 && n < 5 {
		fmt.Print("Inserisci posizione delle carta da tenere: ")
		fmt.Scanln(&c1, &c2, &c3, &c4)
		indici = append(indici, c1, c2, c3, c4)

		drawNewCard := deck.DrawHand(n)

		n = 5 - n
		i := 0
		var newHand []*Card
		for n > 0 {
			n--
			newHand = append(newHand, hand[indici[i]-1])
			i++
		}
		fmt.Print("Carte ricevute: ")
		Print(drawNewCard)
		newHand = append(newHand, drawNewCard...)
		return newHand
	}
	return hand
}

func main() {

	var deck = CreateDeck(9)
	for i := 1; i < 5; i++ {
		hand := deck.DrawHand(5)
		Print(hand)
		changed := (ChangeCards(hand, deck))
		Print(changed)
		fmt.Println(" Punteggio:", CheckScore(changed))

	}
}
