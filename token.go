package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Token struct {
	Name  			string
	X     			int
	Y     			int
	img_file_name	string
	I 	  			*ebiten.Image
}
