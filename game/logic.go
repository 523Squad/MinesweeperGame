package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)


const EasyLvlDimension = 9
const EasyLvlBombsNumber = 10
const MediumLvlDimension = 16
const MediumLvlBombsNumber = 40
const HardLvlDimension = 30
const HardLvlBombsNumber = 116

type point struct {
	isOpened          bool
	hasFlag           bool
	hasMine           bool
	neighborMineCount int
}

// Board represents minesweeper board.
type Board struct {
	bombsNumber int
	dimension   int
	field       [][]*point
	gameOver    bool
}

func (p *point) toString() string {
	return " " + strconv.FormatBool(p.hasMine) + " neighbours " + strconv.Itoa(p.neighborMineCount)
}

func (b *Board) setBoard(n int) {
	for i := 0; i < n; i++ {
		row := []*point{}
		for j := 0; j < n; j++ {
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
		if b.field[x][y].hasMine == false {
			b.field[x][y].hasMine = true
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
					if (ki == 0 && kj == 0) {
						continue
					} else if (((ki + i >= 0) && (ki + i < b.dimension)) &&
						((kj + j >= 0) && (kj + j < b.dimension))) {
						if (b.field[ki + i][kj + j].hasMine == true) {
							b.field[i][j].neighborMineCount++
						}
					}
				}
			}
		}
	}
}

// continue tells whether we should keep playing.
func (b *Board) continuePlaying() bool {
	// Stub. TODO: Implement
	return true
}

func (b *Board) initGame() {
	b.dimension = EASY_LVL_DIM
	b.bombsNumber = EASY_LVL_BOMBS
	b.setBoard(b.dimension)
	fmt.Println(strconv.FormatBool(b.gameOver))

}


