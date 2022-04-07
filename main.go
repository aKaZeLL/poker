package main

import (
	"fmt"
	"math/rand"
	"sort"
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
	seeds := []string{"Cuori", "Quadri", "Picche", "Fiori"}
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

func String(n int) string {
	if n > 1 && n < 11 {
		return fmt.Sprintf("%d", n)
	}
	seeds := map[int]string{11: "J", 12: "Q", 13: "K", 14: "A"}
	return seeds[n]
}

func (d *Deck) DrawHand(n int) []*Card {
	temp := d.Cards[:n]
	d.Cards = d.Cards[n:]
	return temp
}

func Flush(hand []*Card) bool {
	color := hand[0].Seed
	for _, m := range hand {
		if m.Seed != color {
			return false
		}
	}
	return true
}

func StructToArray(hand []*Card) []int {
	var array []int
	for _, h := range hand {
		array = append(array, h.Value)
	}
	return array
}

func Straight(hand []*Card) bool {
	array := StructToArray(hand)
	sort.Ints(array)

	for i := 0; i < len(hand)-1; i++ {
		if array[i] != array[i+1]-1 {
			return false
		}
	}
	return true
}

func Royale(hand []*Card) bool {
	if Straight(hand) && Flush(hand) {
		return true
	} else {
		return false
	}
}

func Count(hand []*Card, q int) int {
	array := StructToArray(hand)

	for i := 0; i < len(array)-1; i++ {
		cnt := 1
		for x := i + 1; x < len(array); x++ {
			if array[i] == array[x] {
				cnt += 1
			}
		}
		if cnt == q {
			return array[i]
		}
	}
	return 0
}

func Poker(hand []*Card) int {
	score := Count(hand, 4)
	if score != 0 {
		return score
	}
	return 0
}

func WithOutElem(wOe []*Card, elem int) []*Card {
	var SubHand []*Card
	for _, value := range wOe {
		if value.Value != elem {
			SubHand = append(SubHand, value)
		}
	}
	return SubHand
}

func Full(hand []*Card) (int, int) {
	tri := Count(hand, 3)
	if tri != 0 {
		if sub := WithOutElem(hand, tri); sub[0].Value == sub[1].Value {
			return tri, sub[0].Value
		}
	}
	return 0, 0
}

func Tris(hand []*Card) int {
	score := Count(hand, 3)
	if score != 0 {
		return score
	}
	return 0
}

func Pair(hand []*Card) int {
	score := Count(hand, 2)
	if score != 0 {
		return score
	}
	return 0
}

func Double(hand []*Card) (int, int) {
	couple := Count(hand, 2)
	if couple != 0 {
		sub := WithOutElem(hand, couple)
		double := Count(sub, 2)
		if double != 0 {
			return couple, double
		}
	}
	return 0, 0
}

func CheckScore(hand []*Card) int {
	if Royale(hand) {
		fmt.Println("Scala Reale!")
		return 10
	} else if s := Poker(hand); s != 0 {
		fmt.Println("Poker di", String(s))
		return 9
	} else if Flush(hand) {
		fmt.Println("Colore")
		return 8
	} else if s, s2 := Full(hand); s != 0 {
		fmt.Println("Full di", String(s), String(s2))
		return 7
	} else if Straight(hand) {
		fmt.Println("Scala Semplice")
		return 6
	} else if s := Tris(hand); s != 0 {
		fmt.Println("Tris di", String(s))
		return 5
	} else if s, s2 := Double(hand); s != 0 {
		fmt.Println("Doppia coppia di", String(s), String(s2))
		return 4
	} else if s := Pair(hand); s != 0 {
		fmt.Println("Coppia di", String(s))
		return 3
	}
	fmt.Println("Carta Alta")
	return 2

}

func Print(hand []*Card) {
	for _, h := range hand {
		fmt.Printf("%3s di %5s", String(h.Value), h.Seed)
	}
	fmt.Println()
}

/*
da implementare
carta alta
*/

func main() {

	var deck = CreateDeck(9)
	for i := 1; i < 5; i++ {
		hand := deck.DrawHand(5)
		Print(hand)
		fmt.Println("Punteggio:", CheckScore(hand))

	}
}
