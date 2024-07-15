package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
)

/*
func main() {
	//var b Board = board_load("01")
	//board_show(b)
}
*/

func show(v interface{}){
	spew.Dump(v)
}

func showstop(v interface{}){
	spew.Dump(v)
	os.Exit(0)
}

func board_show(b Board){
	fmt.Println("Mostrar el tablero ", b.Size_x, "x", b.Size_y)
	var c = Cell{}
	for y:=1; y<=b.Size_y; y++ {
		for x:=1; x<=b.Size_x; x++ {
			c = *b.Cells[x][y]
			fmt.Println("- ", c.Name, c.X, c.Y, c.Mov)
		}
	}
}
