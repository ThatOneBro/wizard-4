package main

import (
	"cart/sprites"

	"github.com/orsinium-labs/wasm4go/w4"
)

type Direction struct {
	X int8
	Y int8
}

type PlayerState int

const (
	Idle PlayerState = iota
)

// String - Creating common behavior - give the type a String function
func (s PlayerState) String() string {
	return [...]string{"Idle"}[s]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (s PlayerState) EnumIndex() int {
	return int(s)
}

type PlayerUpdate struct {
	PlayerVector Vector
}

type Player struct {
	Sprite         *sprites.Sprite
	ScreenPosition w4.Point
	WorldPosition  w4.Point
	State          PlayerState
}

func (p *Player) Update(u *PlayerUpdate) {
	AddVectorToPoint(&p.WorldPosition, &u.PlayerVector)
}

func (p *Player) Draw() {
	wp := GetScreenPosFromWorldPos(p.WorldPosition)
	LerpPointToPoint(&p.ScreenPosition, &wp, 0.2)
	w4.DrawColors.Set(p.Sprite.ColorMapping[0], p.Sprite.ColorMapping[1], p.Sprite.ColorMapping[2], p.Sprite.ColorMapping[3])
	w4.Blit(p.Sprite.GetBytes(), p.ScreenPosition, p.Sprite.Size, p.Sprite.Flags)
}
