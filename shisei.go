package main

import (
	"fmt"
)

const (
	BoardSize = 8
	Empty     = " "
	Black     = "B"
	White     = "W"
)

type Board [BoardSize][BoardSize]string

var currentPlayer string

func main() {
	board := initializeBoard()
	currentPlayer = Black
	fmt.Println("Game over!")
	printBoard(board)
}

func initializeBoard() Board {
	var board Board
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			board[i][j] = Empty
		}
	}

	board[3][3] = White
	board[3][4] = Black
	board[4][3] = Black
	board[4][4] = White

	return board
}

func printBoard(board Board) {
	fmt.Println("  A B C D E F G H")
	for i := 0; i < BoardSize; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < BoardSize; j++ {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
