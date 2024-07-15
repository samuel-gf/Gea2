package main

import (
	_ "bufio"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "golang.org/x/image/font"
	_ "golang.org/x/image/font/opentype"
	_ "image/color"
	"log"
	"net"
	"os"
	_ "strings"
	"time"
)

var img *ebiten.Image
var board Board


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


func init(){
var err error
img, _, err = ebitenutil.NewImageFromFile("img/player2.png")
	if err != nil {
		log.Fatal(err)
	}
}


func (g *Game) Update() error {
	time.Sleep(time.Second)
	return nil
}


func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	screen.DrawImage(img, nil)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 768
}

func NewGameData() GameData {
	g := GameData{
		ScreenWidth:  80,
		ScreenHeight: 50,
		TileWidth:    16,
		TileHeight:   16,
	}
	return g
}

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Hello, World!")

	con, err := net.Dial("tcp", "192.168.1.56:8080")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer con.Close()

	var cmd Command
	fmt.Println("Solicito tablero")
	SendCmd(con, Command{Cmd: "Getboard", Arg: [9]string{"01"}})
	ReceiveCmd(con, &cmd)
	board = Board{Title: cmd.Arg[0], Size_x: cmd.X, Size_y: cmd.Y}
	board.Cells = board_create_cells(cmd.X, cmd.Y)
	fmt.Println("Recibida petición de tablero", 
				board.Title, board.Size_x, board.Size_y)
	// Ahora debo recibir las celdas hasta recibir una orden de stop
	for {
		c := ReceiveCell(con)
		fmt.Println("Recibida celda ", c.X, c.Y, c.Name)
		board.Cells[c.X][c.Y] = &c
		if cmd.Cmd == "Stop" {
			break
		}
		//time.Sleep(time.Second)
	}
	fmt.Println("He recibido un Stop")
	board_show(board)

	

	/*
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
	*/
}
