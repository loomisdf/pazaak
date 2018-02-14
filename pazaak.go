package main

import (
    "fmt"
    "time"
    "strconv"
    "math/rand"
    "errors"
)

type card struct {
    name string
    value int
}

type deck struct {
   name string
   size int
   cards []card
}

func setupDealer() deck {
    _size := 40
    d := deck {
        name: "dealer-deck",
        size: _size,
        cards: make([]card, _size),
    }
    i := 1 // Value for card
    j := 0 // count from 0 - 3
    k := 0 // place in array
    for {
        d.cards[k] = card {
            name: strconv.Itoa(i),
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
    d.cards = append(d.cards[:i], d.cards[i+1:]...)
    return c, nil
}

func playerDeck(name string) deck {
    _size := 10
    d := deck {
        name: name,
        size: _size,
        cards: make([]card, _size),
    }
    d.cards[0] = card { name: "+1", value: 1 }
    d.cards[1] = card { name: "-1", value: -1 }
    d.cards[2] = card { name: "+2", value: 2 }
    d.cards[3] = card { name: "-2", value: -2 }
    d.cards[4] = card { name: "+3", value: 3 }
    d.cards[5] = card { name: "-3", value: -3 }
    d.cards[6] = card { name: "+4", value: 4 }
    d.cards[7] = card { name: "-4", value: -4 }
    d.cards[8] = card { name: "+5", value: 5 }
    d.cards[9] = card { name: "+6", value: 6 }
    return d
}

func main() {
    dealerDeck := setupDealer()

    fmt.Printf("game start\n")
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
