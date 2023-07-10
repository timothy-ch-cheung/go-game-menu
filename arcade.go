package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type direction int

const (
	up    direction = 0
	down  direction = 180
	left  direction = 270
	right direction = 90
	speed           = 4.5
)

type character struct {
	sprite *ebiten.Image
	x      float64
	y      float64
	dir    direction
}

func degToRads(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func (char *character) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(degToRads(float64(char.dir)))
	//op.GeoM.Translate(char.x, char.y)
	x := char.x
	y := char.y

	width := float64(char.sprite.Bounds().Dx() / 2)
	height := float64(char.sprite.Bounds().Dy() / 2)
	switch char.dir {
	case up:
		{
			x -= width
			y -= height
		}
	case down:
		{
			x += width
			y += height
		}
	case left:
		{
			x -= width
			y += height
		}
	case right:
		{
			x += width
			y -= height
		}
	}
	op.GeoM.Translate(x, y)
	screen.DrawImage(char.sprite, op)
}

func (char *character) update() {
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		char.x += speed
		char.dir = right
	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		char.x -= speed
		char.dir = left
	} else if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		char.y -= speed
		char.dir = up
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		char.y += speed
		char.dir = down
	}

}

func createArrow() (*character, error) {
	sprite, err := newImageFromFile("assets/graphics/arrow.png")

	if err != nil {
		return nil, err
	}

	return &character{
		sprite: sprite,
		x:      screenWidth / 2,
		y:      screenHeight / 2,
		dir:    up,
	}, nil
}
