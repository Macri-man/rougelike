package rougelike

import (
	"fmt"
	"github.com/nsf/termbox-go"
)


//function for main loop


func main(){
	
	//set up termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.HideCursor()

	for {

	}
}