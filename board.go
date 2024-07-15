package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	READ_MODE_GENERAL = 1
	READ_MODE_CELLS   = 2
)

type Board struct {
	Title  string
	Size_x int
	Size_y int
	Cells  [][]*Cell
}


/*
func board_load(filename string){
	fmt.Println("HOLA", filename)
}
*/

// Load from a yaml file
func board_load(filename string) Board{
	var b Board
	var c *Cell = &Cell{X: -1, Y: -1}
	// Open the file
	file, err := os.Open("boards/" + filename + ".yaml")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// Read the file, line per line
	lector := bufio.NewScanner(file)
	var read_mode int = READ_MODE_GENERAL
	for lector.Scan() {
		linea := lector.Text()
		sl_fields := strings.Split(linea, ":")
		val := strings.Trim(sl_fields[1], " \n")
		switch read_mode {
		case READ_MODE_GENERAL: // Data of board
			switch strings.Trim(sl_fields[0], " \n") {
			case "title":
				b.Title = val
			case "size_x":
				b.Size_x, _ = strconv.Atoi(val)
			case "size_y":
				b.Size_y, _ = strconv.Atoi(val)
			case "cells":
				read_mode = READ_MODE_CELLS
				b.Cells = board_create_cells(b.Size_x, b.Size_y)
			default: // El nombre del campo es erroneo
				fmt.Fprintf(os.Stderr, "Error en el campo de entrada %s\n", sl_fields[0])
			}

		/* ----- READ_MODE_CELLS ----------------- */
		case READ_MODE_CELLS: // Mode read cells
			if sl_fields[0][2] == '-' { // New register
				sl_fields[0] = sl_fields[0][3:]
				if c.X!=-1 && c.Y!=-1 {
					b.Cells[c.X][c.Y] = c
					c = &Cell{X: -1, Y: -1}
				}
			}
			key := strings.Trim(sl_fields[0], " \n")
			val := strings.Trim(sl_fields[1], " \n")
			switch key {
			case "name":
				c.Name = val
			case "x":
				c.X, _ = strconv.Atoi(val)
			case "y":
				c.Y, _ = strconv.Atoi(val)
			case "mov":
				f64, _ := strconv.ParseFloat(val, 32)
				c.Mov = float32(f64)
			}
		}
	}
	b.Cells[c.X][c.Y] = c
	return b
}

func board_create_cells(x, y int) (c [][]*Cell){
	c = make([][]*Cell, x+1)
	for i := 0; i <= x; i++ { // Create columns
		c[i] = make([]*Cell, y+1)
	}
	if x != 0 && y != 0 {
		for col:=1; col<=y; col++ {
			for fila:=1; fila<=x; fila++ {
				c[fila][col] = &Cell{Name: "...", X: fila, Y:col}
			}
		}
	}
	return c
}
