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

		if board[x][y] != Empty {
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
	_, err := fmt.Sscanf(parts[1], "%d", &y)
	if err != nil {
		return x, y, color
	}
	y-- // 调整为数组索引

	// 解析输入的颜色
	color = strings.ToUpper(parts[2])

	return x, y, color
}

func isValidColor(color string) bool {
	return color == Black || color == White
}

//你可以在这个代码框中输入x、y和color的值，以便放置棋子。例如，输入"A 1 B"表示在A1位置放置黑子。输入"quit"可以退出游戏。
//请在下方输入x、y和color的值，并按Enter键执行放置棋子的操作。このコードボックスにx、y、colorの値を入力して駒を置くことができます。 たとえば、「A 1 B」と入力すると、A1の位置に黒点が配置されます。 「quit」と入力するとゲームを終了できます。
//下にx、y、colorの値を入力し、Enterキーを押して駒を置く操作を実行してください
