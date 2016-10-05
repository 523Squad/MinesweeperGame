package main

import (
	"fmt"
	"minesweeper/game"
	"minesweeper/menu"
)

func main() {
	manager := &menu.Manager{}
	g := &game.Board{}
	if err := manager.Run(g); err != nil {
		fmt.Printf("%v", err)
	}
}
