package game

import (
	"fmt"
	"math/rand"
	"strconv"
)

// TODO: Rename constants.
// see: https://stackoverflow.com/questions/22688906/go-naming-conventions-for-const
const EASY_LVL_DIM = 9
const EASY_LVL_BOMBS = 10
const MIDDLE_LVL_DIM = 16
const MIDDLE_LVL_BOMBS = 40
const HARD_LVL_DIM = 30
const HARD_LVL_BOMBS = 116

type point struct {
	touched     bool
	isBomb      bool
	bombsNumber int
}

// Board represents minesweeper board.
type Board struct {
	bombsNumber int
	dimension   int
	field       [][]*point
	gameOver    bool
}

func (p *point) toString() string {
	return " " + strconv.FormatBool(p.isBomb)
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
}

func (b *Board) setBombs() {
	count := b.bombsNumber
	for count > 0 {
		x := rand.Intn(b.dimension)
		y := rand.Intn(b.dimension)
		if b.field[x][y].isBomb == false {
			b.field[x][y].isBomb = true
			count--
		}
	}
	//c := 0
	//for i := 0; i < b.dimension; i++ {
	//	for j := 0; j < b.dimension; j++ {
	//		if (b.field[i][j].isBomb == true) {
	//			fmt.Println(b.field[i][j].toString() + "		" + strconv.Itoa(c))
	//			c++
	//		} else {
	//			fmt.Println(b.field[i][j].toString())
	//		}
	//
	//
	//	}
	//}
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
