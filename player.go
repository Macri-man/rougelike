package main

import (
	"fmt"
)

type player struct {
	x,y int
	char rune
	hp int
	name string
	weapons map[string]weapon
	armory map[string]armor
	food map[string]consumable
	magic map[string]magicObject
}

type user interface {
	print()
	move()
}