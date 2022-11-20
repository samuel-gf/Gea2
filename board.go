package main

import (
	_ "fmt"
)

type Board struct {
	name   string
	size_x int
	size_y int
	c      [][]Cell
}

func init_board(b Board, size_x int, size_y int) {
	b.size_x = size_x
	b.size_y = size_y
	b.c = make([][]Cell, 10)			// Crea las filas
	for i := 0; i < size_y-1; i++ {		// Crea las columnas
		b.c[i] = make([]Cell, 10)
	}
	for x := 0; x < size_x-1; x++ {
		for y := 0; y < size_y-1; y++ {
			b.c[x][y] = Cell{"Terreno", "", ""}
		}
	}
}
