package main

import (
	"fmt"
)

type small struct {
	x,y int
	char rune
	hp int
	name string
	damage int
	sight int
}

type heavy struct {
	x,y int
	char rune
	hp int
	name string
	damage int
	armor armor
	weapon weapon
	sight int
}

type giant struct {
	x,y int
	char rune
	hp int
	name string
	damage int
	armor armor
	weapon weapon
	sight int
}






