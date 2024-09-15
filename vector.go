package main

import (
	"cart/sprites"

	"github.com/orsinium-labs/wasm4go/w4"
)

type Vector struct {
	X int8
	Y int8
}

var (
	UpVector    = &Vector{X: 0, Y: -1}
	DownVector  = &Vector{X: 0, Y: 1}
	LeftVector  = &Vector{X: -1, Y: 0}
	RightVector = &Vector{X: 1, Y: 0}
)

func AddVectorToPoint(p *w4.Point, v *Vector) {
	p.X = uint8(int8(p.X) + v.X)
	p.Y = uint8(int8(p.Y) + v.Y)
}

func (v *Vector) AddVector(o *Vector) {
	v.X += o.X
	v.Y += o.Y
}

func GetScreenPosFromWorldPos(p w4.Point) w4.Point {
	p.X *= sprites.SpriteWidth
	p.Y *= sprites.SpriteWidth
	return p
}
