package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Carta struct {
	Valore int
	Seme   string
}

func (c Carta) IsFigure() bool {
	switch c.Valore {
	case 1, 11, 12, 13:
		return true
	default:
		return false
	}
}

func (c Carta) ToFig() string {
	semi := map[int]string{1: "A", 11: "J", 12: "Q", 13: "K"}
	return semi[c.Valore]
}

func New(v int, s string) *Carta {
	return &Carta{Valore: v, Seme: s}
}

func CreaMazzo() [52]*Carta {
	colori := []string{"Cuori", "Quadri", "Picche", "Fiori"}
	var mazzo [52]*Carta
	i := 0
	for _, c := range colori {
		for num := 1; num <= 13; num++ {
			mazzo[i] = New(num, c)
			i++
		}
	}
	//Mescola
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(mazzo), func(i, j int) {
		mazzo[i], mazzo[j] = mazzo[j], mazzo[i]
	})
	return mazzo
}

func main() {

	var carte [52]*Carta = CreaMazzo()
	for _, k := range carte {
		if k.IsFigure() {
			fmt.Println(k.ToFig(), k.Seme)
		} else {
			fmt.Println(k.Valore, k.Seme)
		}
	}
}
