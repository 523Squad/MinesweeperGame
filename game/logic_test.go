package game

import (
	"fmt"
	"testing"
)

var board = &Board{}

func initBoard() {
	board = &Board{}

	board.bombsNumber = 2
	board.dimension = 3

	for i := 0; i < board.dimension; i++ {
		row := []*point{}
		for j := 0; j < board.dimension; j++ {
			row = append(row, new(point))
		}
		board.field = append(board.field, row)
	}

	// set bombs
	board.field[0][0].isBomb = true
	board.field[2][1].isBomb = true
	board.flagsLeft = board.bombsNumber
}

func TestInitGame(t *testing.T) {
	board.initGame(EasyMode)

	if board.dimension != EasyLvlDimension ||
		board.bombsNumber != EasyLvlBombsNumber {
		error := fmt.Sprintf("Expected [%d %d], got [%d %d]",
			EasyLvlDimension,
			EasyLvlBombsNumber,
			board.dimension,
			board.bombsNumber)
		t.Error(error)
	}

	board = &Board{}
	board.initGame(MediumMode)

	if board.dimension != MediumLvlDimension ||
		board.bombsNumber != MediumLvlBombsNumber {
		error := fmt.Sprintf("Expected [%d %d], got [%d %d]",
			MediumLvlDimension,
			MediumLvlBombsNumber,
			board.dimension,
			board.bombsNumber)
		t.Error(error)
	}

	board = &Board{}
	board.initGame(HardMode)

	if board.dimension != HardLvlDimension ||
		board.bombsNumber != HardLvlBombsNumber {
		error := fmt.Sprintf("Expected [%d %d], got [%d %d]",
			HardLvlDimension,
			HardLvlBombsNumber,
			board.dimension,
			board.bombsNumber)
		t.Error(error)
	}

	board = &Board{}
}

func TestPointNeighbours(t *testing.T) {
	initBoard()
	board.setBombsNeighbours()
	if board.field[0][1].bombsNumber != 1 {
		error := fmt.Sprintf("[%d %d] Expected %d, got %d", 0, 1,
			1, board.field[0][1].bombsNumber)
		t.Error(error)
	}
	if board.field[1][1].bombsNumber != 2 {
		error := fmt.Sprintf("[%d %d] Expected %d, got %d", 1, 1,
			2, board.field[1][1].bombsNumber)
		t.Error(error)
	}
	if board.field[2][2].bombsNumber != 1 {
		error := fmt.Sprintf("[%d %d] Expected %d, got %d", 2, 2,
			1, board.field[2][2].bombsNumber)
		t.Error(error)
	}
}

func TestShowAllBombs(t *testing.T) {
	initBoard()

	if board.field[0][0].touched != false &&
		board.field[1][2].touched != false {
		t.Error("ShowAllBombsTest error")
	}

	board.showAllBombs()

	if board.field[0][0].touched != true &&
		board.field[1][2].touched != true {
		t.Error("ShowAllBombsTest error")
	}
}

func TestRightClick(t *testing.T) {
	initBoard()
	board.field[0][2].touched = true

	board.flag(0, 2)

	if board.field[0][2].hasFlag != false {
		error := fmt.Sprintf("[%d %d] Expected %t, got %t", 0, 2,
			false, board.field[0][2].hasFlag)
		t.Error(error)
	}

	board.flag(2, 2)

	if board.field[2][2].hasFlag != true {
		error := fmt.Sprintf("[%d %d] Expected %t, got %t", 2, 2,
			true, board.field[2][2].hasFlag)
		t.Error(error)
	}
}

func TestIsWin(t *testing.T) {
	initBoard()
	board.field[0][0].touched = true // one bomb marked by flag, other - closed
	if board.isWin() != false {
		error := fmt.Sprintf("Expected %t, got %t", false, board.isWin())
		t.Error(error)
	}

	for i := 0; i < board.dimension; i++ {
		for j := 0; j < board.dimension; j++ {
			if board.field[i][j].isBomb == false {
				board.field[i][j].touched = true
			}
		}
	}

	if board.isWin() != true {
		error := fmt.Sprintf("Expected %t, got %t", true, board.isWin())
		t.Error(error)
	}
}

func TestResetGame(t *testing.T) {
	initBoard()
	board.gameWin = true

	board.resetGame()
	if board.gameWin != false || board.gameOver != false {
		error := fmt.Sprintf("Expected [%t %t], got [%t %t]",
			false, false, board.gameWin, board.gameOver)
		t.Error(error)
	}
}

func TestBoardChoose(t *testing.T) {
	initBoard()
	board.choose(0, 0)

	if board.gameOver != true {
		error := fmt.Sprintf("[%d %d]Expected %t, got %t",
			0, 0, true, board.gameOver)
		t.Error(error)
	}

	// Open all bombs after game over
	for i := 0; i < board.dimension; i++ {
		for j := 0; j < board.dimension; j++ {
			if board.field[i][j].isBomb &&
				board.field[i][j].touched != true {
				error := fmt.Sprintf("[%d %d]Expected %t, got %t",
					i, j, true, board.gameOver)
				t.Error(error)
			}
		}
	}
}

func TestBoardTouched(t *testing.T) {
	initBoard()
	board.choose(0, 1)

	if board.field[0][1].touched != true {
		error := fmt.Sprintf("[%d %d]Expected %t, got %t",
			0, 1, true, board.field[0][1].touched)
		t.Error(error)
	}
}

func TestPointToString(t *testing.T) {
	p := point{true, true, 4, false}
	s := " true neighbours 4"
	if p.toString() != s {
		error := fmt.Sprintf("Expected %s, got %s", s, p.toString())
		t.Error(error)
	}
}
