package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	playGame()
}

func playGame() {
	board := initializeBoard()
	currentPlayer = Black

	reader := bufio.NewReader(os.Stdin)

	for {
		printBoard(board)
		fmt.Printf("Current Player: %s\n", currentPlayer)

		fmt.Print("Enter x, y, and color (e.g., A 1 B): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			fmt.Println("Invalid input!")
			continue
		}

		x, y, color := parseInput(parts)
		if x < 0 || x >= BoardSize || y < 0 || y >= BoardSize {
			fmt.Println("Invalid coordinates!")
			continue
		}

		if !isValidColor(color) {
			fmt.Println("Invalid color!")
			continue
		}

		if !isEmptyPosition(board, x, y) {
			fmt.Println("The specified position is not empty!")
			continue
		}

		board[x][y] = color

		// 切换玩家
		if currentPlayer == Black {
			currentPlayer = White
		} else {
			currentPlayer = Black
		}
	}

	fmt.Println("Game over!")
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

func parseInput(parts []string) (int, int, string) {
	x := -1
	y := -1
	color := ""

	if len(parts[0]) != 1 || len(parts[1]) != 1 || len(parts[2]) != 1 {
		return x, y, color
	}

	// 解析输入的x坐标
	switch strings.ToUpper(parts[0]) {
	case "A":
		x = 0
	case "B":
		x = 1
	case "C":
		x = 2
	case "D":
		x = 3
	case "E":
		x = 4
	case "F":
		x = 5
	case "G":
		x = 6
	case "H":
		x = 7
	default:
		return x, y, color
	}

	// 解析输入的y坐标
	switch parts[1] {
	case "1":
		y = 7
	case "2":
		y = 6
	case "3":
		y = 5
	case "4":
		y = 4
	case "5":
		y = 3
	case "6":
		y = 2
	case "7":
		y = 1
	case "8":
		y = 0
	default:
		return x, y, color
	}

	// 解析输入的颜色
	switch strings.ToUpper(parts[2]) {
	case "B":
		color = Black
	case "W":
		color = White
	default:
		return x, y, color
	}

	return x, y, color
}

func isValidColor(color string) bool {
	return color == Black || color == White
}

func isEmptyPosition(board Board, x, y int) bool {
	return board[x][y] == Empty
}
