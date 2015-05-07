package rougelike


import (
	"fmt"
)

type player struct {
	x,y int
	hp int
	name string
	armory map[string]armor
	inventory map[string]item
	weapons map[string]weapon
}

type user interface {
	pickItem(state *State)
	releaseItem(state *State)
	useItem(state *State)
	equipeItem(state *State)
}