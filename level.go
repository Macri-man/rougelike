package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type cell int

const (
	ground cell = iota 
	monster
	wall
	empty
)

type room struct {
	x,y,w,h int
	contains []cell
}

type level struct {
	numRooms int
	width, height int
	contains []cell
	errors error
	
}

type tile struct {
	attribute map[string]int 
}

func (c cell) Draw() rune {
	switch c {
	case ground:
		return '.'
	case monster:
		return 'M'
	case empty:
		return ' '
	case wall:
		return '#'
	}
	return '!'
}

func newLevel(w, h int) level {
	l := level{width: w, height: h, contains: make([]cell, w*h)}



	for x := 0; x < w; x++ {
		l.SetCell(x, 0, wall)
		l.SetCell(x, h-1, wall)
	}
	for y := 0; y < h; y++ {
		l.SetCell(0, y, wall)
		l.SetCell(w-1, y, wall)
	}
	return l;
} 


func (l level) At(x, y int) cell {
	return l.contains[x+width*y]
}


func (l *level) SetCell(x, y int, c cell) {
	if x+width*y >= len(l.contains) || x < 0 || y < 0 {
		return
	}
	l.contains[x+width*y] = c
}

func (l level) Draw() {
	for x := 0; x < l.width; x++ {
		for y := 0; y < l.height; y++ {
			termbox.SetCell(x, y, l.At(x, y).Draw(), termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	termbox.Flush()
}

