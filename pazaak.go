package main

import (
	"fmt"
	"os"
)

type GameState struct {
	p1Turn bool
}

var commands = map[string]string{
	"end turn": "end turn",
	"stand":    "stand",
	"quit":     "quit",
}

var gameState GameState
var dealerDeck Deck
var p1, p2 Player

func init() {
	dealerDeck = SetupDealer()
	p1 = Player{
		name:     "player1",
		sidedeck: PlayerDeck("p1Deck"),
		hand:     make([]Card, 4),
		score:    0,
		lastMove: "game start",
	}
	p2 = Player{
		name:     "player2",
		sidedeck: PlayerDeck("p2Deck"),
		hand:     make([]Card, 4),
		score:    0,
		lastMove: "game start",
	}
	gameState = GameState{
		p1Turn: true,
	}
}

// bool returns true if the command is valid
func ParseInput(input string) bool {
	_, ok := commands[input]
	var status bool
	if ok {
		switch input {
		case "end turn":
			status = true
		case "stand":
			status = true
		case "quit":
			os.Exit(0)
		default:
			status = false
		}
	} else {
		status = false
	}
	return status
}

func main() {
	fmt.Printf("game start\n")
	DealHand(&p1)
	DealHand(&p2)
	fmt.Printf("%+v\n", p1)
	fmt.Printf("%+v\n", p2)
	for {
		fmt.Printf("--- Begin turn ---\n")
		c, err := DealCard(&dealerDeck)
		fmt.Printf("dealer deck size = %d\n", len(dealerDeck.Cards))
		if err != nil {
			fmt.Printf("deck is empty\n")
		} else {
			fmt.Printf("%+v\n", c)
		}
		var input string
		fmt.Scanln(&input)
		ParseInput(input)
		fmt.Printf("Here is your input:%s\n", input)
		fmt.Printf("--- End turn ---\n")
	}
	fmt.Printf("game end\n")
}
