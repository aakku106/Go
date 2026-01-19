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

func (p *Position) MoveRight() {
	p.x += 2
}
func (p *Position) MoveLeft() {
	p.x -= 2
}
func (p *Position) MoveUp() {
	p.y += 2
}
func (p *Position) MoveDown() {
	p.y -= 2
}
func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}
func main() {
	aakku := NewPlayer()
	aakku.MoveUp()
	aakku.MoveLeft()
	fmt.Println(aakku.Position)
}
