package main

import (
	"fmt"
)

type Carta struct {
	Valore int
	Seme   string
}

func (c Carta) Value2Color(carta *Carta) string {
	semi := map[int]string{1: "A", 11: "J", 12: "Q", 13: "K"}
	switch carta.Valore {
	case 1, 11, 12, 13:
		return semi[carta.Valore]
	default:
		return carta.Seme
	}
}

func (c Carta) Color2Value(carta *Carta) int {
	valori := map[string]int{"A": 1, "J": 11, "Q": 12, "k": 13}

	switch carta.Seme {
	case "A", "J", "Q", "K":
		return valori[carta.Seme]
	default:
		return carta.Valore
	}
}

func New(v int, s string) *Carta {
	return &Carta{Valore: v, Seme: s}
}

func CreaMazzo() [52]*Carta {
	var colori []string
	var mazzo [52]*Carta
	var i int
	colori = append(colori, "Cuori", "Quadri", "Picche", "Fiori")
	i = 0
	for _, c := range colori {
		for num := 1; num <= 13; num++ {
			mazzo[i] = New(num, c)
			i++
		}
	}
	return mazzo
}

func main() {

	var carte [52]*Carta = CreaMazzo()
	for _, k := range carte {
		fmt.Println(k.Valore, k.Seme)
	}

}
