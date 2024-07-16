package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type User struct {
	Age  int
	Name string
}

func main() {
	//ln, err := net.Listen("tcp", "192.168.1.56:8080")
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer ln.Close()

	fmt.Println("Servidor inicializado. Esperando conexión...")
	for {
		con, err := ln.Accept()
		if err != nil {
			fmt.Println("Server: ", err)
			continue
		}
		fmt.Println("Conexión establecida con cliente")
		go handleClient(con)
		time.Sleep(time.Second) // Duerme un segundo
	}
}

func handleClient(con net.Conn) {
	defer con.Close()
	var cmd Command = Command{}
	for {
		cmd = Command{}
		ReceiveCmd(con, &cmd)
		if cmd.Cmd == "" {
			log.Println("Client disconnected")
			return
		}
		switch cmd.Cmd {
		case "Getboard":
			id := cmd.Arg[0]
			log.Printf("Debo enviar el tablero %s\n", id)
			log.Println("Inicializando mundo ...")
			var board Board = Board{}
			board = board_load(id)
			// Enviar datos del tablero
			fmt.Printf("Tablero del tamaño: %d x %d\n", board.Size_x, board.Size_y)
			Send(con, board)

			// Enviar las celdas
			for y := 1; y <= board.Size_y; y++ {
				for x := 1; x <= board.Size_x; x++ {
					Send(con, board.Cells[x][y])
				}
			}
			Send(con, Command{Cmd: "Stop"})
			log.Println("Tablero enviado al cliente")
		}
	}
}
