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

func printBoard(b Board) {
	fmt.Printf("\n")
	fmt.Printf("| %s | %s | %s |\n", b.top.left, b.top.middle, b.top.right)
	fmt.Printf("| %s | %s | %s |\n", b.middle.left, b.middle.middle, b.middle.right)
	fmt.Printf("| %s | %s | %s |\n", b.bottom.left, b.bottom.middle, b.bottom.right)
	fmt.Printf("\n")
}

func drawShape(spot int, shape string, b Board) Board {
	switch spot {
	case 1:
		if b.top.left == "_" {
			b.top.left = shape
			return b
		} 
	case 2:
		if b.top.middle == "_" {
			b.top.middle = shape
			return b
		}
	case 3:
		if b.top.right == "_" {
			b.top.right = shape
			return b
		}
	case 4:
		if b.middle.left == "_" {
			b.middle.left = shape
			return b
		}
	case 5:
		if b.middle.middle == "_" {
			b.middle.middle = shape
			return b
		}
	case 6:
		if b.middle.right == "_" {
			b.middle.right = shape
			return b
		}
	case 7:
		if b.bottom.left == "_" {
			b.bottom.left = shape
			return b
		}
	case 8:
		if b.bottom.middle == "_" {
			b.bottom.middle = shape
			return b
		}
	}
	fmt.Println("Invalid input")
	return b
}

func gameLoop(b Board) {
	scanner := bufio.NewScanner(os.Stdin)
	printBoard(b)
	someoneWon := false
	for  !someoneWon{
		fmt.Println("Pick a number between 1-9: ")
		scanner.Scan()
		input, err := strconv.Atoi(scanner.Text())
		if err == nil && input >= 1 && input <= 9 {
			b = drawShape(input, "X", b)
			someoneWon = evaluateWin(b, "X") // checks to see if X won the game
			printBoard(b)
			randomInt := rand.Intn(9-1) +1
			fmt.Println("Random Int: ",randomInt)
			b = drawShape(randomInt, "O", b)
		}
	} 
}

func main() {
	fmt.Printf("Tic-Tac-Toe Game\n")
	gameBoard := emptyBoard()
	gameLoop(gameBoard)
}