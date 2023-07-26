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
var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func main() {
	playGame()
}

func playGame() {
	board := initializeBoard()
	currentPlayer = Black

	reader := bufio.NewReader(os.Stdin)

	for {
		printBoard(board)
		if currentPlayer == Black {
			fmt.Printf("Current Player: %s\n", "Black - ○")
		}
		if currentPlayer == White {
			fmt.Printf("Current Player: %s\n", "White - ●")
		}
		// fmt.Printf("Current Player: %s\n", currentPlayer)

		put_count := CheckPut(board, currentPlayer)
		fmt.Print("place to put: ", put_count, "\n")

		// fmt.Print("Enter x, y, and color (e.g., A 1 B): ")
		fmt.Print("Enter x, y (e.g., A 1): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input!")
			continue
		}

		x, y := parseInput(parts)
		if x < 0 || x >= BoardSize || y < 0 || y >= BoardSize {
			fmt.Println("Invalid coordinates!")
			continue
		}

		if !isValidColor(currentPlayer) {
			fmt.Println("Invalid color!")
			continue
		}

		if !isEmptyPosition(board, x, y) {
			fmt.Println("The specified position is not empty!")
			continue
		}

		if !reverse(&board, x, y, currentPlayer) {
			fmt.Println("Don't put here")
			continue
		}

		board[x][y] = currentPlayer

		empty_count := count(board)
		if empty_count == 0 {
			fmt.Println("Finish!!!!!!!!!!!!!!!!!!!!")
			printBoard(board)
			break
		}

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
			if board[i][j] == Empty {
				fmt.Printf("%s ", " ")
			}
			if board[i][j] == Black {
				fmt.Printf("%s ", "○")
			}
			if board[i][j] == White {
				fmt.Printf("%s ", "●")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func parseInput(parts []string) (int, int) {
	x := -1
	y := -1
	// color := ""

	if len(parts[0]) != 1 || len(parts[1]) != 1 {
		return x, y
	}

	// 解析输入的x坐标
	switch strings.ToUpper(parts[0]) {
	case "H":
		y = 7
	case "G":
		y = 6
	case "F":
		y = 5
	case "E":
		y = 4
	case "D":
		y = 3
	case "C":
		y = 2
	case "B":
		y = 1
	case "A":
		y = 0
	default:
		return x, y
	}

	// 解析输入的y坐标（调整映射）
	switch parts[1] {
	case "1":
		x = 0
	case "2":
		x = 1
	case "3":
		x = 2
	case "4":
		x = 3
	case "5":
		x = 4
	case "6":
		x = 5
	case "7":
		x = 6
	case "8":
		x = 7
	default:
		return x, y
	}

	// 解析输入的颜色
	// switch strings.ToUpper(parts[2]) {
	// case "B":
	// 	color = Black
	// case "W":
	// 	color = White
	// default:
	// 	return x, y, color
	// }

	return x, y
}

func isValidColor(color string) bool {
	return color == Black || color == White
}

func isEmptyPosition(board Board, x, y int) bool {
	return board[x][y] == Empty
}

func reverse(board *Board, x, y int, color string) bool {
	check_reverse := false
	for _, dir := range directions {
		flip := [][2]int{}
		i, j := x+dir[0], y+dir[1]

		for i >= 0 && i < 8 && j >= 0 && j < 8 {
			if board[i][j] == Empty {
				break
			}
			if board[i][j] == color {
				for _, pos := range flip {
					board[pos[0]][pos[1]] = color
					check_reverse = true
				}
				break
			}
			flip = append(flip, [2]int{i, j})
			i, j = i+dir[0], j+dir[1]
		}
	}
	return check_reverse
}

func count(board Board) int {
	empty_count := 0
	black_count := 0
	white_count := 0
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if board[i][j] == Empty {
				empty_count += 1
			}
			if board[i][j] == Black {
				black_count += 1
			}
			if board[i][j] == White {
				white_count += 1
			}
		}
	}
	if empty_count == 0 {
		if black_count > white_count {
			fmt.Printf("Winner is Black!! ")
		}
		if black_count < white_count {
			fmt.Printf("Winner is White!! ")
		}
		if black_count == white_count {
			fmt.Printf("This game is draw.")
		}
	}
	if empty_count != 0 {
		fmt.Printf("------------------------------------------\n")
		fmt.Printf("Black=%d, White=%d \n", black_count, white_count)
	}
	return empty_count
}

func CheckPut(board Board, currentPlayer string) int {
	put_count := 0
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if board[i][j] == Empty {
				if reverse(&board, i, j, currentPlayer) {
					put_count += 1
				}
			}
		}
	}
	return put_count
}
