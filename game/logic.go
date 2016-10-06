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

func (b *Board) setBombsNeighbours() {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			coords := []int{-1, 0, 1}
			for _, ki := range coords {
				for _, kj := range coords {
					if ki == 0 && kj == 0 {
						continue
					} else if ((ki+i >= 0) && (ki+i < b.dimension)) &&
						((kj+j >= 0) && (kj+j < b.dimension)) {
						if b.field[ki+i][kj+j].isBomb == true {
							b.field[i][j].bombsNumber++
						}
					}
				}
			}
		}
	}
}

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
	b.updateState(newBoardState)
}

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
				b.updateState(newBoardState)
				if b.isWin() {
					b.gameWin = true
				}
			} else {
				coords := []int{-1, 0, 1}
				for _, ki := range coords {
					for _, kj := range coords {
						if ki == 0 && kj == 0 {
							continue
						} else {
							nextI, nextJ := ki+row, kj+col
							if isCoordValid(nextI, nextJ, b.dimension) && !newBoardState[nextI][nextJ].hasFlag {
								b.choose(nextI, nextJ)
							}
						}
					}
				}
			}
		}
	}
}

func isCoordValid(i, j, dim int) bool {
	return i >= 0 && j >= 0 && i < dim && j < dim
}

func (b *Board) updateState(newBoard [][]*point) {
	b.field = newBoard
}

func (b *Board) isWin() bool {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if !b.field[i][j].isBomb && !b.field[i][j].touched {
				return false
			}
		}
	}
	return true
}

// continuePlaying tells whether we should keep playing.
func (b *Board) continuePlaying() bool {
	return !b.gameOver
}

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

func (b *Board) showAllBombs() {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.field[i][j].isBomb {
				b.field[i][j].touched = true
			}
		}
	}
}

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

func (b *Board) resetGame() {
	b.field = [][]*point{}
	b.gameWin = false
	b.gameOver = false
}

//func main() {
//	b := Board {dimension:EasyLvlDimension,
//		   bombsNumber: EasyLvlBombsNumber, }
//	b.setBoard(b.dimension)
//	fmt.Println(strconv.FormatBool(b.gameOver))
//	b.showBoard()
//	fmt.Println()
//	fmt.Println()
//	count := 0
//	for b.continuePlaying() {
//		b.performLeftClick(rand.Intn(b.dimension),rand.Intn(b.dimension))
//		b.showBoard()
//		fmt.Println()
//		fmt.Println()
//		count++
//	}
//	fmt.Println(count)
//}
