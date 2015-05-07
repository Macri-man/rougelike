package rougelike


import (
	"fmt"
)

type item interface {
	action
	user
}

type magicObject struct {
	x,y int
	item
	level int
	name string
}

type weapon struct {
	x,y int 
	item
	level int 
	damage int 
	special int
	name string
}

type cosumable struct {
	x,y int 
	item
	level int 
	special int
	name string
}

type action interface {
	
}
