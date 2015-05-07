package rougelike

import (
	"fmt"
	"errors"

)


type level struct {
	rooms []room
	errors error
}

type tile struct {
	x,y int
	char byte
	attribute map[int]string 
	creature creature
}