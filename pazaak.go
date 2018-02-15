package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type card struct {
	name  string
	value int
}

type deck struct {
	name  string
	size  int
	cards []card
}

type player struct {
	name     string
	sidedeck deck
	hand     []card
	score    int
}

func setupDealer() deck {
	_size := 40
	d := deck{
		name:  "dealer-deck",
		size:  _size,
		cards: make([]card, _size),
	}
	i := 1 // Value for card
	j := 0 // count from 0 - 3
	k := 0 // place in array
	for {
		d.cards[k] = card{
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

func dealCard(d *deck) (card, error) {
	if d.size == 0 {
		return card{}, errors.New("deck is empty")
	}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	i := r.Intn(d.size)
	c := d.cards[i]
	d.size--
	d.cards = append(d.cards[:i], d.cards[i+1:]...) // remove the card
	return c, nil
}

func playerDeck(name string) deck {
	_size := 10
	d := deck{
		name:  name,
		size:  _size,
		cards: make([]card, _size),
	}
	d.cards[0] = card{name: "+1", value: 1}
	d.cards[1] = card{name: "-1", value: -1}
	d.cards[2] = card{name: "+2", value: 2}
	d.cards[3] = card{name: "-2", value: -2}
	d.cards[4] = card{name: "+3", value: 3}
	d.cards[5] = card{name: "-3", value: -3}
	d.cards[6] = card{name: "+4", value: 4}
	d.cards[7] = card{name: "-4", value: -4}
	d.cards[8] = card{name: "+5", value: 5}
	d.cards[9] = card{name: "+6", value: 6}
	return d
}

func dealHand(p *player) {
	size := 10
	c := make([]card, size) // Make a copy of the players deck
	copy(c, p.sidedeck.cards)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < 4; i++ {
		pos := r.Intn(size) // random card from deck
		p.hand[i] = c[pos]
		c = append(c[:pos], c[pos+1:]...) // remove the card
		size--
	}
}

func main() {
	dealerDeck := setupDealer()
	p1 := player{
		name:     "player1",
		sidedeck: playerDeck("p1Deck"),
		hand:     make([]card, 4),
		score:    0,
	}
	p2 := player{
		name:     "player2",
		sidedeck: playerDeck("p2Deck"),
		hand:     make([]card, 4),
		score:    0,
	}

	fmt.Printf("game start\n")
	dealHand(&p1)
	dealHand(&p2)
	fmt.Printf("%+v\n", p1)
	fmt.Printf("%+v\n", p2)
	for {
		fmt.Printf("--- Begin turn ---\n")
		c, err := dealCard(&dealerDeck)
		fmt.Printf("dealer deck size = %d\n", len(dealerDeck.cards))
		if err != nil {
			fmt.Printf("deck is empty\n")
		} else {
			fmt.Printf("%+v\n", c)
		}
		fmt.Scanln()
		fmt.Printf("--- End turn ---\n")
	}
	fmt.Printf("game end\n")
}
