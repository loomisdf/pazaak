package main

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

type Card struct {
	name  string
	value int
}

type Deck struct {
	name  string
	size  int
	Cards []Card
}

type Player struct {
	name     string
	sidedeck Deck
	hand     []Card
	score    int
	lastMove string
}

func SetupDealer() Deck {
	_size := 40
	d := Deck{
		name:  "dealer-deck",
		size:  _size,
		Cards: make([]Card, _size),
	}
	i := 1 // Value for Card
	j := 0 // count from 0 - 3
	k := 0 // place in array
	for {
		d.Cards[k] = Card{
			name:  strconv.Itoa(i),
			value: i,
		}
		if i == 10 && j == 3 {
			break
		}
		if j == 3 {
			i++
			j = 0
		} else {
			j++
		}
		k++
	}
	return d
}

func DealCard(d *Deck) (Card, error) {
	if d.size == 0 {
		return Card{}, errors.New("deck is empty")
	}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	i := r.Intn(d.size)
	c := d.Cards[i]
	d.size--
	d.Cards = append(d.Cards[:i], d.Cards[i+1:]...) // remove the Card
	return c, nil
}

func PlayerDeck(name string) Deck {
	_size := 10
	d := Deck{
		name:  name,
		size:  _size,
		Cards: make([]Card, _size),
	}
	d.Cards[0] = Card{name: "+1", value: 1}
	d.Cards[1] = Card{name: "-1", value: -1}
	d.Cards[2] = Card{name: "+2", value: 2}
	d.Cards[3] = Card{name: "-2", value: -2}
	d.Cards[4] = Card{name: "+3", value: 3}
	d.Cards[5] = Card{name: "-3", value: -3}
	d.Cards[6] = Card{name: "+4", value: 4}
	d.Cards[7] = Card{name: "-4", value: -4}
	d.Cards[8] = Card{name: "+5", value: 5}
	d.Cards[9] = Card{name: "+6", value: 6}
	return d
}

func DealHand(p *Player) {
	size := 10
	c := make([]Card, size) // Make a copy of the players deck
	copy(c, p.sidedeck.Cards)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < 4; i++ {
		pos := r.Intn(size) // random Card from deck
		p.hand[i] = c[pos]
		c = append(c[:pos], c[pos+1:]...) // remove the Card
		size--
	}
}
