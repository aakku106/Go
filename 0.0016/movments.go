package main

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
