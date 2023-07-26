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
	"io/ioutil"
)

const (
	BoardSize  = 8
	SquareSize = 60
	LineWidth  = 2.0
	Empty     = "0"
        Black     = "B"
        White     = "W"
)

var (
	BoardColors = []string{"#FFCE9E", "#D18B47"}
	PieceColors = []string{"#000000", "#FFFFFF"} // black and white
)

type Board [BoardSize][BoardSize]string

var currentPlayer string

var color string

var is_load bool

var filename string

var orderfile string

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	currentPlayer = Black
	filename = "saved_game.txt"
	orderfile = "player_order.txt"
 
	fmt.Print("Do you wanna load game?（true or false）: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	_, err := fmt.Sscan(scanner.Text(), &is_load)
	if err != nil {
		fmt.Println("input error：", err)
		return
	}

	var board Board

	if is_load {
		fmt.Println("game was loaded")
		board, err = loadBoardFromTxt(filename)
		if err != nil {
			fmt.Println("load failed:", err)
			return
		}
		currentPlayer, err = loadCurrentPlayer(orderfile)
    		if err != nil {
        		fmt.Println("Error loading currentPlayer:", err)
       			return
    		}
	} else {
		board = initializeBoard()
		fmt.Println("start a new game")
	}

	clearInputBuffer(scanner)

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
                	fmt.Println("image load failed:", err)
        	}

		fmt.Printf("Current Player: %s\n", currentPlayer)

		fmt.Print("Enter x, y(e.g., A 1) save/quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			break
		}

		if input == "save"{
			err := saveBoardToTxt(board, filename)
			if err != nil {
				fmt.Println("save failed:", err)
				return
			}
		        err = saveCurrentPlayer(currentPlayer, orderfile)
   			if err != nil {
        			fmt.Println("Error saving currentPlayer:", err)
        			return
    			}
			fmt.Println("game saved")
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


		if !isEmptyPosition(board, x, y) {
			fmt.Println("The specified position is not empty!")
			continue
		}

		color := currentPlayer
		board[x][y] = color
		reverse(&board, x, y, color)

		empty_count := count(board)
		if empty_count == 0 {
			fmt.Println("Finish!!!!!!!!!!!!!!!!!!!!")
			break
		}

		// switch player
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
        // open image by command "xdg-open"
        cmd := exec.Command("xdg-open", filename)
        err := cmd.Start()
        if err != nil {
                return err
        }

        // close image
        time.Sleep(1 * time.Second)

        // kill process
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
			if board[i][j] == "0"{
				fmt.Printf("  ")
			}else{
				fmt.Printf("%s ", board[i][j])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func parseInput(parts []string) (int, int) {
	x := -1
	y := -1

	if len(parts[0]) != 1 || len(parts[1]) != 1{
		return x, y
	}

	// x axis
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

	// y axis
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


	return x, y
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

func saveBoardToTxt(board Board, filename string) error {
        data := make([]string, BoardSize)
        for i := 0; i < BoardSize; i++ {
                data[i] = strings.Join(board[i][:], " ")
        }
        content := strings.Join(data, "\n")

        err := ioutil.WriteFile(filename, []byte(content), 0644)
        if err != nil {
                return err
        }
        return nil
}

func loadBoardFromTxt(filename string) (Board, error) {
        var board Board

        content, err := ioutil.ReadFile(filename)
        if err != nil {
                return board, err
        }

        lines := strings.Split(string(content), "\n")
        for i, line := range lines {
                data := strings.Split(line, " ")
                for j, cell := range data {
                        board[i][j] = cell
                }
        }
        return board, nil
}


func clearInputBuffer(scanner *bufio.Scanner) {
	reader := bufio.NewReader(os.Stdin)
	reader.Discard(reader.Buffered())
}

func saveCurrentPlayer(currentPlayer string, filename string) error {
    return ioutil.WriteFile(filename, []byte(currentPlayer), 0644)
}

func loadCurrentPlayer(filename string) (string, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", err
    }
    return string(data), nil
}
