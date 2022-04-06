package main

import (
	"fmt"
	"math/rand"
	"time"
	"sort"
)

type Card struct {
	Value int
	Seed   string
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

func (c Card) String() string {
	if c.Value > 1 && c.Value < 11 {
		return fmt.Sprintf("%3d di %5s", c.Value, c.Seed)
	}
	seeds := map[int]string{11: "J", 12: "Q", 13: "K", 14: "A"}
	return fmt.Sprintf("%2s di %5s", seeds[c.Value], c.Seed)
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
	
	for i:=0; i<len(hand)-1; i++ {
		if array[i]!=array[i+1]-1 {
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

func count(hand []*Card, q int) int {
	array := StructToArray(hand)

	for i:=0; i<len(array)-1; i++ {
		cnt := 1
		for x:=i+1; x<len(array); i++ {
			if array[i]==array[x] {
				cnt+=1
			}
		}
		if cnt==q {
			return array[i]
			}
	}
	return 0
}

func Print(hand []*Card) {
	for _, h:= range hand {
		fmt.Print(h.String())
	}
	fmt.Println()
}

/*
reale OK
poker
colore OK
full
scala semplice OK
tris
doppia coppia
coppia
carta alta
*/

func main() {

	var deck = CreateDeck(7)
//	for i:=1; i<5; i++ {
	hand := deck.DrawHand(5)

	Print(hand)


//}
}
