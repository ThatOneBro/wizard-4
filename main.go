package main

import (
	"cart/sprites"

	"github.com/orsinium-labs/wasm4go/w4"
)

var (
	player = &Player{
		Sprite:    &sprites.WizardSprite,
		Position:  w4.Point{X: 2, Y: 0},
		Direction: Direction{X: 1, Y: 0},
	}
	frameCount = 0
)

func init() {
	w4.Start = start
	w4.Update = update
}

func main() {}

func start() {
	w4.Palette.Set(
		w4.Color{R: 0x9b, G: 0xbc, B: 0x0f},
		w4.Color{R: 0x8b, G: 0xac, B: 0x0f},
		w4.Color{R: 0x30, G: 0x62, B: 0x30},
		w4.Color{R: 0x0f, G: 0x38, B: 0x0f},
	)
}

func update() {
	frameCount++

	g := w4.Gamepad

	if g.Up() {
		player.Up()
	} else if g.Down() {
		player.Down()
	} else if g.Left() {
		player.Left()
	} else if g.Right() {
		player.Right()
	} else {
		player.Direction = Direction{}
	}

	if frameCount%15 == 0 {
		player.Update()
	}
	player.Draw()
}
