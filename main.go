package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)


const (
	width  = 80
	height = 24
)

func main(){
	
	//set up termbox
	fmt.Printf("START")
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.HideCursor()
	level := newLevel(80,24)
	for {
		level.Draw()
		event := termbox.PollEvent()
		switch event.Ch {
		case 'w':
		case 'a':
		case 's':
		case 'd':

		case 'q':
			break;
		//updateMap()
		}
	}
}