package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

const EASY_LVL_DIM = 9
const EASY_LVL_BOMBS = 10
const MIDDLE_LVL_DIM = 16
const MIDDLE_LVL_BOMBS = 40
const HARD_LVL_DIM = 30
const HARD_LVL_BOMBS = 116

type Point struct {
	touched     bool
	isBomb      bool
	bombsNumber int
}

type Board struct {
	bombsNumber int
	dimension   int
	field       [][]*Point
	gameOver    bool
}

func (p *Point) toString() string {
	return " " + strconv.FormatBool(p.isBomb)
}

func (b *Board) setBoard(n int) {
	for i := 0; i < n; i++ {
		row := []*Point{}
		for j := 0; j < n; j++ {
			row = append(row, new(Point))
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

func main() {
	p := Board{
		dimension:   EASY_LVL_DIM,
		bombsNumber: EASY_LVL_BOMBS,
	}
	p.setBoard(p.dimension)
	fmt.Println(strconv.FormatBool(p.gameOver))

}
