package main

import (
	"cart/sprites"

	"github.com/orsinium-labs/wasm4go/w4"
)

var (
	player = &Player{
		Sprite:         &sprites.WizardSprite,
		WorldPosition:  w4.Point{X: 0, Y: 0},
		ScreenPosition: w4.Point{X: 0 /* 2 x 16 */, Y: 0},
	}
	frameCount = 0
)

// Need a main, even if its a noop. Otherwise runtime crashes due to undefined symbol
func main() {}

func init() {
	w4.Start = start
	w4.Update = update
}

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
	v := Vector{X: 0, Y: 0}

	if frameCount%15 == 0 {
		if g.Up() {
			v.AddVector(&UpVector)
		}
		if g.Down() {
			v.AddVector(&DownVector)
		}
		if g.Left() {
			v.AddVector(&LeftVector)
		}
		if g.Right() {
			v.AddVector(&RightVector)
		}

		player.Update(&PlayerUpdate{PlayerVector: v})
	}

	player.Draw()
}
