package main

import (
	"fmt"
)

type Position struct {
	x float64
	y float64
}
type Player struct {
	*Position
}
type Enemy struct {
	*Position
}

func main() {}
