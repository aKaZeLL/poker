package valutation

import (
	"fmt"
	"sort"
)

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

func Straight(hand []*Card) int {
	array := StructToArray(hand)
	sort.Ints(array)

	for i := 0; i < len(hand)-1; i++ {
		if array[i] != array[i+1]-1 {
			return 0
		}
	}
	return array[len(array)-1]
}

func Royale(hand []*Card) bool {
	if Straight(hand) != 0 && Flush(hand) {
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
			if couple > double {
				return couple, double
			} else {
				return double, couple
			}
		}
	}
	return 0, 0
}

func CheckScore(hand []*Card) int {
	if Royale(hand) {
		fmt.Println("Scala Reale!")
		return 200
	} else if s := Poker(hand); s != 0 {
		fmt.Println("Poker di", String(s))
		return 150 + s
	} else if Flush(hand) {
		fmt.Println("Colore")
		return 120
	} else if s, s2 := Full(hand); s != 0 {
		fmt.Println("Full di", String(s), String(s2))
		return 90 + s + s2
	} else if s := Straight(hand); s != 0 {
		fmt.Println("Scala Semplice")
		return 75 + s
	} else if s := Tris(hand); s != 0 {
		fmt.Println("Tris di", String(s))
		return 60 + s
	} else if s, s2 := Double(hand); s != 0 {
		fmt.Println("Doppia coppia di", String(s), String(s2))
		return 30 + s + s2
	} else if s := Pair(hand); s != 0 {
		fmt.Println("Coppia di", String(s))
		return 15 + s
	}
	fmt.Println("Carta Alta")
	return 2
}
