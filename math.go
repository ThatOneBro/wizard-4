package main

import "github.com/orsinium-labs/wasm4go/w4"

func Lerp(a, b, t float64) float64 {
	return a + (b-a)*t
}

func LerpPointToPoint(a *w4.Point, b *w4.Point, t int) {
	a.X = uint8(Lerp(float64(a.X), float64(b.X), float64(t)))
	a.Y = uint8(Lerp(float64(a.Y), float64(b.Y), float64(t)))
}
