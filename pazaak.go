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
	winner string
}

var commands = map[string]string{
	"end turn": "end turn",
	"stand":    "stand",
	"quit":     "quit",
}

var gameState GameState
var dealerDeck Deck
var p1, p2 Player

const maxScore = 20
const gameOver = "game over"
const stand = "stand"
const endTurn = "end turn"
const quit = "quit"
const tieGame = "tie game"
const p1Wins = "player 1 wins!"
const p2Wins = "player 2 wins!"

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
		winner: "",
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
		case endTurn:
			status = true
		case stand:
			status = true
		case quit:
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
	if gameState.p1Turn && p1.lastMove != stand {
		// Player 1's turn
		num, err := TakeCard(&dealerDeck, &p1)
		if err != nil {
			fmt.Printf("dealer's deck is empty\n")
		} else {
			fmt.Printf("P1 drew %d, score is %d\n", num, p1.score)
			if p1.score == maxScore {
				p1.lastMove = stand
			} else {
				fmt.Printf("P1, what are you going to do?('end turn', 'stand', or 'quit')\n")
				input := ReadInput()
				for !ParseInput(input) {
					fmt.Printf("invalid command\n")
					input = ReadInput()
				}
				fmt.Printf("p1 turn over\n")
				p1.lastMove = input
			}
		}
	} else if gameState.p1Turn == false && p2.lastMove != stand {
		// Player 2's turn
		num, err := TakeCard(&dealerDeck, &p2)
		if err != nil {
			fmt.Printf("dealer's deck is empty\n")
		} else {
			fmt.Printf("P2 drew %d, score is %d\n", num, p2.score)
			if p2.score == maxScore {
				p2.lastMove = stand
			} else {
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
	}
	if gameState.p1Turn {
		gameState.p1Turn = false
	} else {
		gameState.p1Turn = true
	}
}

func evalGameState() {
	if p1.score > maxScore {
		gameState.winner = p2Wins
		gameState.state = gameOver
	} else if p2.score > maxScore {
		gameState.winner = p1Wins
		gameState.state = gameOver
	} else if p1.lastMove == stand && p2.lastMove == stand {
		if p1.score == p2.score {
			gameState.winner = tieGame
		} else if p1.score > p2.score {
			gameState.winner = p1Wins
		} else {
			gameState.winner = p2Wins
		}
		gameState.state = gameOver
	}
}

func main() {
	fmt.Printf("game start\n")
	for {
		fmt.Printf("--- Begin turn ---\n")
		fmt.Printf("p1 score: %d --- p2 score: %d\n", p1.score, p2.score)
		turn()
		evalGameState()
		if gameState.state == gameOver {
			fmt.Printf(gameState.winner)
			break
		}
		fmt.Printf("--- End turn ---\n")
	}
}
