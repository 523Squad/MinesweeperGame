package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const EasyMode = 0
const MediumMode = 1
const HardMode = 2
const EasyLvlDimension = 9
const EasyLvlBombsNumber = 10
const MediumLvlDimension = 16
const MediumLvlBombsNumber = 40
const HardLvlDimension = 30
const HardLvlBombsNumber = 116

type point struct {
	touched     bool
	isBomb      bool
	bombsNumber int
	hasFlag     bool
}

// Board represents minesweeper board.
type Board struct {
	bombsNumber int
	flagsLeft   int
	dimension   int
	field       [][]*point
	gameOver    bool
	gameWin     bool
}

func (p *point) toString() string {
	return " " + strconv.FormatBool(p.isBomb) + " neighbours " + strconv.Itoa(p.bombsNumber)
}

//initialize neccessary params for game : bombs and neighbours
func (b *Board) setBoard() {
	for i := 0; i < b.dimension; i++ {
		row := []*point{}
		for j := 0; j < b.dimension; j++ {
			row = append(row, new(point))
		}
		b.field = append(b.field, row)
	}
	b.setBombs()
	b.setBombsNeighbours()
}

//random bombs generations on the board
func (b *Board) setBombs() {
	rand.Seed(time.Now().UTC().UnixNano())
	count := b.bombsNumber
	for count > 0 {
		x := rand.Intn(b.dimension)
		y := rand.Intn(b.dimension)
		if b.field[x][y].isBomb == false {
			b.field[x][y].isBomb = true
			count--
		}
	}
}

//definiton dangerous neighbors for each point
func (b *Board) setBombsNeighbours() {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			neighbours := getNeighbours(i, j)
			for k := 0; k < 8; k++ {
				nextI := neighbours[k][0]
				nextJ := neighbours[k][1]
				if isCoordValid(nextI, nextJ, b.dimension) {
					if b.field[nextI][nextJ].isBomb == true {
						b.field[i][j].bombsNumber++
					}
				}
			}
		}
	}
}

//perform user's  assumption about bomb location
func (b *Board) flag(row int, col int) {
	newBoardState := b.field
	p := newBoardState[row][col]
	if !p.touched {
		if p.hasFlag {
			p.hasFlag = false
			b.flagsLeft++
		} else if b.flagsLeft > 0 {
			p.hasFlag = true
			b.flagsLeft--
		}
	}
	if b.isWin(newBoardState) {
		b.gameWin = true
	}
	b.updateState(newBoardState)
}

//perform point research
func (b *Board) choose(row int, col int) {
	newBoardState := b.field
	p := newBoardState[row][col]
	if !p.touched {
		if p.hasFlag {
			p.hasFlag = false
		}
		p.touched = true
		if p.isBomb {
			b.updateState(newBoardState)
			b.gameOver = true
			b.showAllBombs()
		} else {
			bombs := newBoardState[row][col].bombsNumber
			if bombs > 0 {
				if b.isWin(newBoardState) {
					b.gameWin = true
				}
				b.updateState(newBoardState)
			} else {
				neighbours := getNeighbours(row, col)
				for i := 0; i < 8; i++ {
					nextI := neighbours[i][0]
					nextJ := neighbours[i][1]
					if isCoordValid(nextI, nextJ, b.dimension) && !newBoardState[nextI][nextJ].hasFlag {
						b.choose(nextI, nextJ)
					}
				}
			}
		}
	}
}

//avoid IndexOutOfBoundException
func isCoordValid(i, j, dim int) bool {
	return i >= 0 && j >= 0 && i < dim && j < dim
}

//slice of point's neighbors
func getNeighbours(row int, col int) [][]int {
	neighbours := make([][]int, 8)
	coords := []int{-1, 0, 1}
	i := 0
	for _, ki := range coords {
		for _, kj := range coords {
			if ki == 0 && kj == 0 {
				continue
			} else {
				nextI, nextJ := ki+row, kj+col
				neighbours[i] = make([]int, 2)
				neighbours[i][0], neighbours[i][1] = nextI, nextJ
				i++
			}
		}
	}
	return neighbours
}

//update board field after user action
func (b *Board) updateState(newBoard [][]*point) {
	b.field = newBoard
}

func (b *Board) isWin(state [][]*point) bool {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if !state[i][j].isBomb && !state[i][j].touched {
				return false
			}
		}
	}
	return true
}

// continuePlaying tells whether we should keep playing.
func (b *Board) continuePlaying() bool {
	return (!b.gameOver && !b.gameWin)
}

//helper for demonstrating board state in console mode
func (b *Board) showBoard() {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.field[i][j].touched {
				if b.field[i][j].isBomb {
					fmt.Print("x" + " ")
				} else {
					fmt.Print(strconv.Itoa(b.field[i][j].bombsNumber) + " ")
				}
			} else {
				fmt.Print("*" + " ")
			}
		}
		fmt.Println()
	}
}

//shows bombs location after game over
func (b *Board) showAllBombs() {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.field[i][j].isBomb {
				b.field[i][j].touched = true
			}
		}
	}
}

//game start with mode according to menu option selected
func (b *Board) initGame(mode int) {
	dimension := -1
	bombsNumber := -1
	switch mode {
	case EasyMode:
		dimension = EasyLvlDimension
		bombsNumber = EasyLvlBombsNumber
	case MediumMode:
		dimension = MediumLvlDimension
		bombsNumber = MediumLvlBombsNumber
	case HardMode:
		dimension = HardLvlDimension
		bombsNumber = HardLvlBombsNumber
	}
	b.dimension = dimension
	b.bombsNumber = bombsNumber
	b.flagsLeft = bombsNumber
	b.setBoard()
}

//reset board params to default for the new game
func (b *Board) resetGame() {
	b.field = [][]*point{}
	b.gameWin = false
	b.gameOver = false
}
