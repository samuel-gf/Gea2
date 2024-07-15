package main

/* Comandos:
	Getboard 01
	Setboard
	Setcell
	Stop
*/

import (
	"net"
	"encoding/json"
	"log"
	"os"
	"fmt"
	"bufio"
	"time"
)

type Command struct {
	Cmd string
	Arg [9]string
	X, Y, Z int
}

func SendCmd(con net.Conn, cmd Command){
	bytes, err := json.Marshal(cmd)
	if err != nil {
		log.Println("Error json.Marshal")
		os.Exit(1)
	}
	fmt.Fprintf(con, string(bytes)+"\n")
	time.Sleep(time.Second/50)
}

func Send(con net.Conn, data interface{}){
	bytes, err := json.Marshal(data)
	if err != nil{
		log.Println("Error json.Marshal")
		os.Exit(1)
	}
	fmt.Fprintf(con, string(bytes)+"\n")
	time.Sleep(time.Second/100)
}

func ReceiveCmd(con net.Conn, cmd *Command) {
	reader := bufio.NewReader(con)
	str, _ := reader.ReadString('\n')
	if len(str) > 0 {
		err := json.Unmarshal([]byte(str), cmd)
		if err != nil {
			log.Println("Error unmarshal")
			log.Println(err)
			os.Exit(2)
		}
	}
}

func ReceiveCell(con net.Conn) (c Cell) {
	reader := bufio.NewReader(con)
	str, _ := reader.ReadString('\n')
	if len(str) > 0 {
		err := json.Unmarshal([]byte(str), &c)
		if err != nil {
			log.Println("Error unmarshal")
			log.Println(err)
			os.Exit(2)
		}
	}
	return c
}
