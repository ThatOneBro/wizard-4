package sprites

import "github.com/orsinium-labs/wasm4go/w4"

var WizardSprite = &Sprite{
	Bytes:        [64]byte{0x00, 0x00, 0xf0, 0x00, 0x00, 0x03, 0xa8, 0x00, 0x00, 0x0e, 0xac, 0x00, 0x00, 0x0e, 0xac, 0x00, 0x00, 0x38, 0x0b, 0x00, 0x00, 0xea, 0xaa, 0xc0, 0x03, 0xaa, 0xaa, 0xb0, 0x0e, 0xe7, 0x36, 0xc0, 0x08, 0xb0, 0x43, 0x00, 0x02, 0xa8, 0x0a, 0xc0, 0x02, 0x29, 0x5a, 0xc0, 0x03, 0xb9, 0x5b, 0x00, 0x00, 0xb9, 0x5b, 0x00, 0x02, 0xaf, 0xfe, 0x80, 0x00, 0xaa, 0xaa, 0x00, 0x00, 0x00, 0x00, 0x00},
	ColorMapping: [4]w4.DrawColor{w4.Light, w4.Primary, w4.Secondary, w4.Dark},
	Flags:        w4.TwoBPP,
	Size:         w4.Size{Width: 16, Height: 16},
}
