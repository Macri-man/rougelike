package main

import (
	"fmt"
)

type item interface {
	print() 
}

type magicObject struct {
	level int
	char rune
	healthmod int
	special int
	name string
}

type weapon struct {
	level int 
	char rune
	healthmod int
	durability int
	special int
	name string
}

type armor struct {
	level int 
	char rune
	healthmod int
	durability int
	special int
	name string
}

type consumable struct {
	char rune
	healthmod int
	duration int
	special int
	name string
}

func (c consumable) print() {
	fmt.Printf("%+v\n", c)
}

func (a armor) print() {
	fmt.Printf("%+v\n", a)
}

func (m magicObject) print() {
	fmt.Printf("%+v\n", m)
}

func (w weapon) print() {
	fmt.Printf("%+v\n", w)	
}

func printItem(i item) {
	i.print()
}
