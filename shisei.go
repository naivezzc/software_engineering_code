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

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

type Board [BoardSize][BoardSize]string

var currentPlayer string

func main() {
	board := initializeBoard()
	//currentPlayer = Black
	//fmt.Println("Game over!")
	//printBoard(board)
	fmt.Println("Before:")
	printBoard(board)

	board[2][3] = "B"
	reverse(&board, 2, 3, "B")

	fmt.Println("\nAfter:")
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

func reverse(board *Board, x, y int, color string) {
	for _, dir := range directions {
		flip := [] [2]int{}
		i, j := x+dir[0], y+dir[1]

		for i >= 0 && i < 8 && j >= 0 && j < 8 {
			if board[i][j] == " " {
				break
			}
			if board[i][j] == color {
				for _, pos := range flip {
					board[pos[0]][pos[1]] = color
				}
				break
			}
			flip = append(flip, [2]int{i, j})
			i, j = i+dir[0], j+dir[1]
		}
	}
}
