package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type GameState struct {
	p1Turn bool
	state  string
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
	reset()
}

func reset() {
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
	DealHand(&p1)
	DealHand(&p2)
	gameState = GameState{
		p1Turn: true,
		state:  "game start",
	}
}

func ReadInput() string {
	var reader = bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = regexp.MustCompile(`\r?\n`).ReplaceAllString(input, "") // remove newlines with support for windows
	return input
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

func turn() {
	if gameState.p1Turn && p1.lastMove != "stand" {
		// Player 1's turn
		num, err := TakeCard(&dealerDeck, &p1)
		if err != nil {
			fmt.Printf("dealer's deck is empty\n")
		} else {
			fmt.Printf("P1 drew %d, score is %d\n", num, p1.score)
			fmt.Printf("P1, what are you going to do?('end turn', 'stand', or 'quit')\n")
			input := ReadInput()
			for !ParseInput(input) {
				fmt.Printf("invalid command\n")
				input = ReadInput()
			}
			fmt.Printf("p1 turn over\n")
			p1.lastMove = input
		}
	} else if gameState.p1Turn == false && p2.lastMove != "stand" {
		// Player 2's turn
		num, err := TakeCard(&dealerDeck, &p2)
		if err != nil {
			fmt.Printf("dealer's deck is empty\n")
		} else {
			fmt.Printf("P2 drew %d, score is %d\n", num, p2.score)
			fmt.Printf("P2, what are you going to do?('end turn', 'stand', or 'quit')\n")
			input := ReadInput()
			for !ParseInput(input) {
				fmt.Printf("invalid command\n")
				input = ReadInput()
			}
			fmt.Printf("p2 turn over\n")
			p2.lastMove = input
		}
	}
	if gameState.p1Turn {
		gameState.p1Turn = false
	} else {
		gameState.p1Turn = true
	}
}

func main() {
	fmt.Printf("game start\n")
	for {
		fmt.Printf("--- Begin turn ---\n")
		fmt.Printf("p1 score: %d --- p2 score: %d\n", p1.score, p2.score)
		turn()
		if gameState.state == "game over" {
			break
		}
		fmt.Printf("--- End turn ---\n")
	}
	fmt.Printf("game end\n")
}
