package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/hajimehoshi/ebiten/v2"
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
	Bg_file_name string
	Bg  *ebiten.Image
	//Tokens [10]Token
}

// Load from a yaml file
func board_load(filename string) Board {
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
			case "bg":
				b.Bg_file_name = val
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
				if c.X != -1 && c.Y != -1 {
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
			case "i":
				c.ImageId = val
			}
		}
	}
	b.Cells[c.X][c.Y] = c
	board_show(b)
	return b
}

// Se usa para rellenar un tablero recien creado
func board_create_cells(x, y int) (c [][]*Cell) {
	c = make([][]*Cell, x+1)
	for i := 0; i <= x; i++ { // Create columns
		c[i] = make([]*Cell, y+1)
	}
	if x != 0 && y != 0 {
		for col := 1; col <= y; col++ {
			for fila := 1; fila <= x; fila++ {
				c[fila][col] = &Cell{Name: "no_name", X: fila, Y: col, Mov: 100, ImageId:"no_image_id"}
			}
		}
	}
	return c
}

func board_show(b Board) {
	fmt.Println("Mostrar el tablero ", b.Size_x, "x", b.Size_y, "Fondo:",b.Bg_file_name)
	var c = Cell{}
	for y := 1; y <= b.Size_y; y++ {
		for x := 1; x <= b.Size_x; x++ {
			c = *b.Cells[x][y]
			fmt.Println("- ", c.Name, c.X, c.Y, c.Mov, c.ImageId)
		}
	}
}

func board_show2(b Board) {
	fmt.Println("MOSTREMOS")
	board_show(b)
	for y := 1; y <= b.Size_y; y++ {
		for x := 1; x <= b.Size_x; x++ {
			if b.Cells[x][y].Mov < 100 {
				fmt.Print(" ")
			} else {
				fmt.Print("█")
			}
		}
		fmt.Println()
	}
}
