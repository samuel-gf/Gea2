package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"net"
	"os"
	"time"
	"math"
	"image/color"
	"strconv"
)

var board Board
var img *ebiten.Image
var tokens [10]Token

type Game struct{}

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

// One tile
type MapTile struct {
	PixelX  int
	PixelY  int
	Blocked bool
	Image   *ebiten.Image
}

func init() {
	var err error
	// Cargar las imágenes
	img, _, err = ebitenutil.NewImageFromFile("img/player2.png")
	if err != nil {
		log.Fatal(err)
	}
	//arr_token[1] = Token{Name:"PJ", X:2, Y:2, I:img[0]}
}

func (g *Game) Update() error {
	time.Sleep(time.Second/10)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Fondo
	screen.DrawImage(board.Bg, nil)
	//var t Token
	for i:=0; i<10; i++{
		var t Token = tokens[i]
		if t.I != nil{
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(70*(t.X-1)+1), float64(70*(t.Y-1)+1-5))
			screen.DrawImage(t.I, op)
		}
	}
	// Casillas
	/*
	for y:=1; y<=board.Size_y; y++{
		for x:=1; x<=board.Size_x; x++{
			//t = board.Tokens[i]
			var c Cell = *board.Cells[x][y]
			//fmt.Println("Dibujar celda ", c.Name, "en", c.X, c.Y)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(70*(c.X-1)+1), float64(70*(c.Y-1)+1))
			if (board.Cells[x][y].I != nil) {
				screen.DrawImage(board.Cells[x][y].I, op)

			}
		}
	}
	*/
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return board.Size_x*70, board.Size_y*70
}


func main() {
	con, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer con.Close()

	fmt.Println("Solicito tablero")
	SendCmd(con, Command{Cmd: "Getboard", Arg: [9]string{"01"}})
	board = ReceiveBoard(con)
	board.Cells = board_create_cells(board.Size_x, board.Size_y)
	fmt.Println("Recibida petición de tablero", board.Title, board.Size_x, board.Size_y)
	// Ahora debo recibir las celdas hasta recibir una orden de stop
	for {
		c := ReceiveCell(con)
		if (c.X == 0 && c.Y == 0) || (c.ImageId == ""){ 		// No se recibirán más celdas
			break
		} else {						// Sí hay celdas para consumir
			//fmt.Println("Cargando imagen", c.ImageId)
			// Cargar la imagen que corresponda a cada celda
			c.I, _, err = ebitenutil.NewImageFromFile("img/"+c.ImageId+".png")
			if err != nil {
				log.Fatal(err)
			}
			board.Cells[c.X][c.Y] = &c
		}
		//time.Sleep(time.Second)
	}
	board.Bg, _, err = ebitenutil.NewImageFromFile("img/bg/"+board.Bg_file_name+".png")
	if err != nil {
		log.Fatal(err)
	}
	// Tokens
	for i:=0; i<4; i++{
		var t *Token = &tokens[i]
		// Face del token
		var face *ebiten.Image
		face, _, err = ebitenutil.NewImageFromFile("img/tokens/00"+strconv.Itoa(i+1)+".png")

		if err != nil {
			log.Fatal(err)
		}
		// Fondo del token
		t.I = ebiten.NewImage(70, 70)
		t.I.Fill(color.RGBA{255, 255, 255, 255})
		t.I.DrawImage(face, nil)
		t.X = 4+i
		t.Y = 4
		var cx int = 35
		var cy int = 35
		var d float64
		for y:=0; y<70; y++{
			for x:=0; x<70; x++{
				d = math.Sqrt((float64)((cx-x)*(cx-x) + (cy-y)*(cy-y)))
				if d > 34 {
					t.I.Set(x, y, color.Transparent)
				}
				if d > 30 && d<34 {
					t.I.Set(x, y, color.RGBA{0, 255, 0, 255})
				}
			}
		}
	}
	tokens[1].Y=3
	tokens[1].X=13
	tokens[3].X=12
	tokens[3].Y=6
	tokens[2].X=14
	tokens[2].Y=9


	// monstruos
	for i:=4; i<=7; i++{
		var t *Token = &tokens[i]
		// Face del token
		var face *ebiten.Image
		face, _, err = ebitenutil.NewImageFromFile("img/tokens/m001.png")
		if err != nil {
			log.Fatal(err)
		}
		// Fondo del token
		t.I = ebiten.NewImage(70, 70)
		t.I.Fill(color.RGBA{255, 255, 255, 255})
		t.I.DrawImage(face, nil)
		t.X = 3
		t.Y = 10
		var cx int = 35
		var cy int = 35
		var d float64
		for y:=0; y<70; y++{
			for x:=0; x<70; x++{
				d = math.Sqrt((float64)((cx-x)*(cx-x) + (cy-y)*(cy-y)))
				if d > 34 {
					t.I.Set(x, y, color.Transparent)
				}
				if d > 30 && d<34 {
					t.I.Set(x, y, color.RGBA{255, 0, 0, 255})
				}
			}
		}
	}
	tokens[5].X++
	tokens[5].Y++
	tokens[6].X+=2
	tokens[7].Y-=2
	tokens[7].X+=3

	board_show2(board)



	//ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(70*board.Size_x, 70*board.Size_y)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Gea 2.0")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
