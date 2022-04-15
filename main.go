package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Player struct {
	Name      string
	Hand      []*Card
	Value     map[string]int
	PointName string
}

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

//Costruttore giocatore
func CreatePlayer(name string, h []*Card) *Player {
	v := map[string]int{"Total": 0, "BestCard": 0, "Best2nd": 0, "Kicker": 0}
	return &Player{Name: name, Hand: h, Value: v}
}

//metodi
func (p Player) Draw(other Player) Player {
	for {
		elem := p.Value["BestCard"]
		p.Hand = WithOutElem(p.Hand, elem)
		other.Hand = WithOutElem(other.Hand, elem)
		if len(p.Hand) == 0 {
			break
		}

		p.Value["BestCard"] = HighCard(p.Hand)
		p.PointName = "Kicker " + String(p.Value["BestCard"])
		other.Value["BestCard"] = HighCard(other.Hand)
		other.PointName = "Kicker " + String(other.Value["BestCard"])

		if winner := p.Valutation(other, "BestCard"); winner.Name != "pareggio" {
			return winner
		}
	}
	return Player{Name: "pareggio"}
}

func (d *Deck) DrawHand(n int) []*Card {
	temp := d.Cards[:n]
	d.Cards = d.Cards[n:]
	return temp
}

//cambio carte mano
func (p *Player) ChangeCards(deck *Deck) {
	var n int
	var c1, c2, c3, c4 int
	var indici []int
	// ERRORE QUI SCOpe VARIabile n
	for n := -1; n < 0 || n >= 5; {
		fmt.Print(p.Name, " quante carte vuoi cambiare: ")
		fmt.Scanln(&n)
		fmt.Println(n)
	}
	fmt.Print(n)
	if n != 0 {
		fmt.Print("Inserisci posizione delle carta da tenere: ")
		fmt.Scanln(&c1, &c2, &c3, &c4)
		indici = append(indici, c1, c2, c3, c4)

		drawNewCard := deck.DrawHand(n)

		n = 5 - n
		i := 0
		var newHand []*Card
		for n > 0 {
			n--
			newHand = append(newHand, p.Hand[indici[i]-1])
			i++
		}
		fmt.Print("Carte ricevute: ")
		for _, v := range drawNewCard {
			fmt.Printf("%s%s ", String(v.Value), v.Seed)
		}
		fmt.Println()
		newHand = append(newHand, drawNewCard...)
		p.Hand = newHand
	}
}

//stampa dati struttura in output
func (p Player) Print() {
	fmt.Print("Mano di ", p.Name, ": ")
	for _, h := range p.Hand {
		fmt.Printf("%s%s ", String(h.Value), h.Seed)
	}
	fmt.Println()
}

//funzioni di supporto
func String(n int) string {
	if n > 1 && n < 11 {
		return fmt.Sprintf("%d", n)
	}
	seeds := map[int]string{11: "J", 12: "Q", 13: "K", 14: "A"}
	return seeds[n]
}

func StructToArray(hand []*Card) []int {
	var array []int
	for _, h := range hand {
		array = append(array, h.Value)
	}
	return array
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

func WithOutElem(wOe []*Card, elem int) []*Card {
	var SubHand []*Card
	for _, value := range wOe {
		if value.Value != elem {
			SubHand = append(SubHand, value)
		}
	}
	return SubHand
}

func MaxInArray(array []int) int {
	max := 0
	for _, i := range array {
		if i > max {
			max = i
		}
	}
	return max
}

//Funzioni valutazione punteggio mano
func Flush(hand []*Card) int {
	color := hand[0].Seed
	for _, m := range hand {
		if m.Seed != color {
			return 0
		}
	}
	array := StructToArray(hand)
	sort.Ints(array)
	return array[len(array)-1]
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

func Royale(hand []*Card) int {
	if v := Straight(hand); (v != 0) && Flush(hand) != 0 {
		return v
	} else {
		return 0
	}
}

func Poker(hand []*Card) int {
	score := Count(hand, 4)
	if score != 0 {
		return score
	}
	return 0
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

func Double(hand []*Card) (int, int, int) {
	couple := Count(hand, 2)
	if couple != 0 {
		sub := WithOutElem(hand, couple)
		double := Count(sub, 2)
		if double != 0 {
			kicker := StructToArray(WithOutElem(hand, double))
			if couple > double {
				return couple, double, kicker[0]
			} else {
				return double, couple, kicker[0]
			}
		}
	}
	return 0, 0, 0
}

func HighCard(hand []*Card) int {
	array := StructToArray(hand)
	return MaxInArray(array)
}

//Valutazione finale punteggio
func (p *Player) CheckScore() {
	v := p.Value
	if v["BestCard"] = Royale(p.Hand); v["BestCard"] != 0 {
		p.PointName = "Scala Reale"
		v["Total"] = 9
	} else if v["BestCard"] = Poker(p.Hand); v["BestCard"] != 0 {
		p.PointName = "Poker di " + String(v["BestCard"])
		v["Total"] = 8
	} else if v["BestCard"] = Flush(p.Hand); v["BestCard"] != 0 {
		p.PointName = "Colore a " + p.Hand[0].Seed
		v["Total"] = 7
	} else if v["BestCard"], v["Best2nd"] = Full(p.Hand); v["BestCard"] != 0 {
		p.PointName = "Full di " + String(v["BestCard"]) + String(v["Best2nd"])
		v["Total"] = 6
	} else if v["BestCard"] = Straight(p.Hand); v["BestCard"] != 0 {
		p.PointName = "Scala Semplice"
		v["Total"] = 5
	} else if v["BestCard"] = Tris(p.Hand); v["BestCard"] != 0 {
		p.PointName = "Tris di " + String(v["BestCard"])
		v["Total"] = 4
	} else if v["BestCard"], v["Best2nd"], v["Kicker"] = Double(p.Hand); v["BestCard"] != 0 {
		p.PointName = "Doppia coppia di " + String(v["BestCard"]) + String(v["Best2nd"])
		v["Total"] = 3
	} else if v["BestCard"] = Pair(p.Hand); v["BestCard"] != 0 {
		p.PointName = "Coppia di " + String(v["BestCard"])
		v["Total"] = 2
	} else {
		v["BestCard"] = HighCard(p.Hand)
		p.PointName = "Carta Alta " + String(v["BestCard"])
		v["Total"] = 1
	}
}

func (p Player) ShoWinner() {
	fmt.Println("Vince ", p.Name, "con ", p.PointName)
}

func (p Player) Valutation(other Player, param string) Player {
	if p.Value[param] > other.Value[param] {
		return p
	} else if other.Value[param] > p.Value[param] {
		return other
	}
	return Player{Name: "pareggio"}
}

func Confront(p1, p2 Player) {
	if winner := p1.Valutation(p2, "Total"); winner.Name != "pareggio" {
		winner.ShoWinner()
	} else if winner := p1.Valutation(p2, "BestCard"); winner.Name != "pareggio" {
		winner.ShoWinner()
	} else {
		//unico caso doppia coppia
		if (p1.Value["Kicker"] != 0) && (p1.Value["Best2nd"] == p2.Value["Best2nd"]) {
			//si valuta kicker
			winner := p1.Valutation(p2, "Kicker")
			winner.ShoWinner()
		} else if p1.Value["Kicker"] != 0 {
			winner := p1.Valutation(p2, "Best2nd")
			winner.ShoWinner()
		} //casi di carta alta e coppia e colore con pareggi
		if p1.Value["Total"] == 1 || p1.Value["Total"] == 2 || p1.Value["Total"] == 7 {
			//carta alta e coppia
			winner := p1.Draw(p2)
			winner.ShoWinner()
		} else {
			//scala semplice\reale pattano, full poker tris non possono pareggiare
			fmt.Println("Partita Patta")
		}
	}
}

func main() {
	//crea il deck con carte mescolate e i giocatori facendoli pescare
	var deck = CreateDeck(8)
	var p1 = CreatePlayer("Fabio", deck.DrawHand(5))
	var p2 = CreatePlayer("Fabrizio", deck.DrawHand(5))

	//TESTING
	//"♥", "♦", "♠", "♣"
	// var slicehand1 []*Card
	// var slicehand2 []*Card
	// c1 := New(8, "♦")
	// c2 := New(4, "♦")
	// c3 := New(7, "♦")
	// c4 := New(12, "♦")
	// c5 := New(14, "♦")
	// slicehand1 = append(slicehand1, c1, c2, c3, c4, c5)

	// a1 := New(8, "♣")
	// a2 := New(3, "♣")
	// a3 := New(7, "♣")
	// a4 := New(12, "♣")
	// a5 := New(14, "♣")
	// slicehand2 = append(slicehand2, a1, a2, a3, a4, a5)

	// var p1 = CreatePlayer("Fabio", slicehand1)
	// var p2 = CreatePlayer("Fabri", slicehand2)
	//FINE TEST

	//stampa i valori
	p1.Print()
	p2.Print()
	//cambio delle carte
	p1.ChangeCards(deck)
	p2.ChangeCards(deck)
	//stampa le carte aggiornate con i cambi
	p1.Print()
	p2.Print()
	//valutazione punteggi
	p1.CheckScore()
	p2.CheckScore()

	Confront(*p1, *p2)
}
