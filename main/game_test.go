package main

import (
	"testing"
)

type Test struct {
	x, y int
}

//testしたい座標を複数用意
var Tests = []Test{
	Test{1, 1},
	// Test{3, 3},
	Test{2, 2},
}

func TestInitializeBoard(t *testing.T) {
	board := initializeBoard()
	if board[3][3] != White {
		t.Errorf("initializeBoard")
	}
	if board[4][4] != White {
		t.Errorf("initializeBoard")
	}
	if board[3][4] != Black {
		t.Errorf("initializeBoard")
	}
	if board[4][3] != Black {
		t.Errorf("initializeBoard")
	}

}
func TestIsEmptyPosition(t *testing.T) {
	board := initializeBoard()
	for _, i := range Tests {
		v := isEmptyPosition(board, i.x, i.y)
		if !v {
			t.Errorf("X = %d, y =  %d is not empty!.", i.x, i.y)
		}
	}
}
