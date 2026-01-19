package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

type Game struct{}

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

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

var aakku = NewPlayer()

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		aakku.MoveUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		aakku.MoveDown()
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		aakku.MoveLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		aakku.MoveRight()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255}) // clear screen to black

	var path vector.Path
	path.MoveTo(float32(aakku.Position.x), float32(aakku.Position.y))
	path.ArcTo(float32(aakku.Position.x), float32(aakku.Position.y), 20, 0, 2*3.14159) // full circle
	path.Close()

	op := &ebiten.DrawTrianglesOptions{}
	vector.Fill(screen, path, color.RGBA{255, 0, 0, 255}, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Size of the logical screen.
	return 640, 480
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Weeeee")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
