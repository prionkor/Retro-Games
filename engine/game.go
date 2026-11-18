package engine

import "github.com/prionkor/retro-games/platform"

type Game interface {
	Parts() []*FramePart
	Render(*Frame)
	Update(dt float64, input platform.Input)
}
