package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Card represents a single playing card.
type Card struct {
	name  string
	value int
}

// Deck is a representation of a game play deck.
//It consists of a set of cards and is used for
//both the dealer and the players.
type Deck struct {
	name  string
	size  int
	Cards []Card
}

// Player represents a single player in the game of pazaak.
type Player struct {
	name     string
	sidedeck Deck
	hand     []Card
	score    int
	lastMove string
}

// SetupDealer returns the dealer deck with cards numbered 1-10 with 4 of each number.
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

// DealCard take a Deck d and removes a random card and returns it.
// DealCard returns an error if the deck is empty.
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

// TakeCard take a deck and player and simulates the player taking a card from the deck and consuming the point value
func TakeCard(d *Deck, p *Player) (int, error) {
	c, err := DealCard(d)
	if err != nil {
		return 0, errors.New("deck is empty")
	}
	p.score += c.value
	return c.value, nil
}

// PlayerDeck returns a new deck with 10 cards.
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

// PrintHand prints the players hand to std out.
func PrintHand(p *Player) {
	fmt.Printf("hand: %s, %s, %s, %s\n",
		p.hand[0].name,
		p.hand[1].name,
		p.hand[2].name,
		p.hand[3].name)
}

// DealHand taks a player p and creates a hand with 4 cards from the player's sidedeck.
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
