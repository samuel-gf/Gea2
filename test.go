package main

import (
	_ "fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
)

/*
func main() {
	//var b Board = board_load("01")
	//board_show(b)
}
*/

func show(v interface{}) {
	spew.Dump(v)
}

func showstop(v interface{}) {
	spew.Dump(v)
	os.Exit(0)
}
