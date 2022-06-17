// tic-tac-toe game implemented in GO for practice learning GO


// get input from command line
// print tik-tak-toe board

package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"math/rand"
)

var SCANNER = bufio.NewScanner(os.Stdin) // global scanner
var globalBoard = emptyBoard() // global board

type Row struct {
	left string
	middle string
	right string	
}

type Board struct {
	top Row
	middle Row
	bottom Row
}

func emptyRow() Row {
	return Row{"_", "_", "_"}
}

func emptyBoard() Board {
	return Board{emptyRow(), emptyRow(), emptyRow()}
}

func printBoard() {
	fmt.Printf("\n")
	fmt.Printf("| %s | %s | %s |\n", globalBoard.top.left, globalBoard.top.middle, globalBoard.top.right)
	fmt.Printf("| %s | %s | %s |\n", globalBoard.middle.left, globalBoard.middle.middle, globalBoard.middle.right)
	fmt.Printf("| %s | %s | %s |\n", globalBoard.bottom.left, globalBoard.bottom.middle, globalBoard.bottom.right)
	fmt.Printf("\n")
}

func drawShape(spot int, shape string) bool{
	switch spot {
	case 1:
		if globalBoard.top.left == "_" {
			globalBoard.top.left = shape
			return true
		} 
	case 2:
		if globalBoard.top.middle == "_" {
			globalBoard.top.middle = shape
			return true
		}
	case 3:
		if globalBoard.top.right == "_" {
			globalBoard.top.right = shape
			return true
		}
	case 4:
		if globalBoard.middle.left == "_" {
			globalBoard.middle.left = shape
			return true
		}
	case 5:
		if globalBoard.middle.middle == "_" {
			globalBoard.middle.middle = shape
			return true
		}
	case 6:
		if globalBoard.middle.right == "_" {
			globalBoard.middle.right = shape
			return true
		}
	case 7:
		if globalBoard.bottom.left == "_" {
			globalBoard.bottom.left = shape
			return true
		}
	case 8:
		if globalBoard.bottom.middle == "_" {
			globalBoard.bottom.middle = shape
			return true 
		}
	case 9:
		if globalBoard.bottom.right == "_" {
			globalBoard.bottom.right = shape
			return true
		}
	}
	return false
}

func playerTurn(shape string) {
	fmt.Printf("Player Turn\nPick a number between 1-9: ")
	SCANNER.Scan()
	input, err := strconv.Atoi(SCANNER.Text()) // gets input from the user and converts it to an int
	if err == nil && input >= 1 && input <= 9 {
		result := drawShape(input, shape)
		if !result {
			fmt.Println("Invalid Input, try again")
			playerTurn(shape)
			return
		}
	}
}
func opponentTurn(shape string) {
	randomInt := rand.Intn(9-1) +1 // a random integer from 1-9
	result := drawShape(randomInt, shape)
	if !result {
		opponentTurn(shape)
	}
}

func evaluateWin(c string) bool {
	return false
}

func gameLoop() {
	printBoard()
	// someoneWon := false
	for  i:=1; i <= 9; i++ {
		if i % 2 != 0 {
			playerTurn("X")
			// someoneWon = evaluateWin(b, "X") // checks to see if X won the game
			printBoard()
		} else {
			opponentTurn("O")
			// someoneWon = evaluateWin(b, "X") // checks to see if X won the game
			fmt.Println("Opponent Turn")
			printBoard()
		}
	}
}

func main() {
	fmt.Printf("Tic-Tac-Toe Game\n")
	gameLoop()
}