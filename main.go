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

func New(v int, s string) *Carta {
	return &Carta{Valore: v, Seme: s}
}

func (c Carta) String() string {
	if c.Valore == 1 {
		return fmt.Sprintf("A %s", c.Seme)
	}

	if c.Valore < 11 {
		return fmt.Sprintf("%d %s", c.Valore, c.Seme)
	}

	semi := map[int]string{1: "A", 11: "J", 12: "Q", 13: "K"}
	return fmt.Sprintf("%s %s", semi[c.Valore], c.Seme)
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
	var carte = CreaMazzo()
	for _, k := range carte {
		fmt.Println(k.String())
	}
}
