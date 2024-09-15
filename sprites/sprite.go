package sprites

import "github.com/orsinium-labs/wasm4go/w4"

const (
	SpriteWidth uint8 = 16
)

type Sprite struct {
	Bytes        [64]byte
	ColorMapping [4]w4.DrawColor
	Flags        w4.BlitFlags
	Size         w4.Size
}

func (s *Sprite) GetBytes() []byte {
	return s.Bytes[:]
}
