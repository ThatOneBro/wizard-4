package main

import (
	"cart/sprites"

	"github.com/orsinium-labs/wasm4go/w4"
)

type Direction struct {
	X int8
	Y int8
}

type Player struct {
	Sprite    *sprites.Sprite
	Position  w4.Point
	Direction Direction
}

func (p *Player) Update() {
	AddDirection(&p.Position, &p.Direction)
}

func (p *Player) Draw() {
	w4.DrawColors.Set(p.Sprite.ColorMapping[0], p.Sprite.ColorMapping[1], p.Sprite.ColorMapping[2], p.Sprite.ColorMapping[3])
	w4.Blit(p.Sprite.GetBytes(), p.Position, p.Sprite.Size, p.Sprite.Flags)
}

func (p *Player) Down() {
	if p.Direction.Y == 0 {
		p.Direction = Direction{X: 0, Y: 1}
	}
}

func (p *Player) Up() {
	if p.Direction.Y == 0 {
		p.Direction = Direction{X: 0, Y: -1}
	}
}

func (p *Player) Left() {
	if p.Direction.X == 0 {
		p.Direction = Direction{X: -1, Y: 0}
	}
}

func (p *Player) Right() {
	if p.Direction.X == 0 {
		p.Direction = Direction{X: 1, Y: 0}
	}
}
