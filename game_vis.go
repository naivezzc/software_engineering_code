package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"math/rand"
        "os/exec"
        "time"
	"bufio"
	"os"
	"strings"
)

const (
	BoardSize  = 8
	SquareSize = 60
	LineWidth  = 2.0
	Empty     = " "
        Black     = "B"
        White     = "W"
)

var (
	BoardColors = []string{"#FFCE9E", "#D18B47"}
	PieceColors = []string{"#000000", "#FFFFFF"} // 黑色和白色
)

type Board [BoardSize][BoardSize]string

var currentPlayer string

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func main() {
	board := initializeBoard()
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)


	for {
		printBoard(board)
		dc := gg.NewContext(BoardSize*SquareSize, BoardSize*SquareSize)
 	 	drawBoard(dc)
	        randomlyPlacePieces(dc, board)
	        err := dc.SavePNG("chessboard.png")
	        
		if err != nil {
                	fmt.Println("Error saving PNG:", err)
        	}

        	imageFiles := []string{"chessboard.png"}
        	
		if err := showImage(imageFiles[0]); err != nil {
                	fmt.Println("无法展示图片:", err)
        	}

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
		reverse(&board, x, y, color)

		empty_count := count(board)
		if empty_count == 0 {
			fmt.Println("Finish!!!!!!!!!!!!!!!!!!!!")
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

func drawBoard(dc *gg.Context) {
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			colorIndex := (i + j) % 2
			dc.SetHexColor(BoardColors[colorIndex])
			dc.DrawRectangle(float64(j*SquareSize), float64(i*SquareSize), float64(SquareSize), float64(SquareSize))
			dc.Fill()
		}
	}

	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(LineWidth)
	for i := 0; i <= BoardSize; i++ {
		dc.DrawLine(0, float64(i*SquareSize), float64(BoardSize*SquareSize), float64(i*SquareSize))
		dc.Stroke()

		dc.DrawLine(float64(i*SquareSize), 0, float64(i*SquareSize), float64(BoardSize*SquareSize))
		dc.Stroke()
	}
}

func randomlyPlacePieces(dc *gg.Context, board Board) {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			if board[i][j] == White {
				colorIndex := 1
				dc.SetHexColor(PieceColors[colorIndex])
                        	dc.DrawCircle(float64(j*SquareSize+SquareSize/2), float64(i*SquareSize+SquareSize/2), SquareSize/2-LineWidth)
                        	dc.Fill()
			} else if board[i][j] == Black{
				colorIndex := 0
				dc.SetHexColor(PieceColors[colorIndex])
	                        dc.DrawCircle(float64(j*SquareSize+SquareSize/2), float64(i*SquareSize+SquareSize/2), SquareSize/2-LineWidth)
        	                dc.Fill()

			} else{
				continue
			}
		}
	}
}


func showImage(filename string) error {
        // 调用系统的默认图片查看器打开图片
        cmd := exec.Command("xdg-open", filename)
        err := cmd.Start()
        if err != nil {
                return err
        }

        // 等待一段时间后关闭图片查看器
        time.Sleep(1 * time.Second)

        // 关闭图片查看器进程
        err = cmd.Process.Kill()
        if err != nil {
                return err
        }

        return nil
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
		return x, y, color
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

func reverse(board *Board, x, y int, color string) {
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
				}
				break
			}
			flip = append(flip, [2]int{i, j})
			i, j = i+dir[0], j+dir[1]
		}
	}
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
