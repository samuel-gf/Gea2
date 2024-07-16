package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Cell struct {
	Name  	string
	Desc  	string
	ImageId	string
	X     	int
	Y     	int
	Mov   	float32
	I	 	*ebiten.Image
}
