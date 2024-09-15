package main

import "github.com/orsinium-labs/wasm4go/w4"

func AddDirection(p *w4.Point, d *Direction) {
	p.X = uint8(int8(p.X) + d.X)
	p.Y = uint8(int8(p.Y) + d.Y)
}
