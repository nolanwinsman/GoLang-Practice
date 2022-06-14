// tic-tac-toe game implemented in GO for practice learning GO


// get input from command line
// print tik-tak-toe board

package main

import "fmt"

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

func main() {
	fmt.Printf("Tic-Tac-Toe Game\n")
	gameBoard := emptyBoard()
	printBoard(gameBoard)
	gameBoard.top.left = "X"
	printBoard(gameBoard)
}