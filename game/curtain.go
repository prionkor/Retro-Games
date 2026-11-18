package game

import "github.com/prionkor/retro-games/engine"

type CurtainState int

const (
	CurtainIdle CurtainState = iota
	CurtainUp
	CurtainDown
)

type CurtainEvent int

const (
	CurtainEventNone CurtainEvent = iota
	CurtainEventCovered
	CurtainEventComplete
)

type Curtain struct {
	Y     float64
	Speed float64
	State CurtainState
}

func NewCurtain(speed float64) *Curtain {
	return &Curtain{
		Y:     60,
		Speed: speed,
		State: CurtainIdle,
	}
}

func (c *Curtain) Trigger() {
	c.Y = 60
	c.State = CurtainUp
}

func (c *Curtain) Animate(dt float64) CurtainEvent {
	switch c.State {
	case CurtainUp:
		c.Y -= c.Speed * dt

		if c.Y <= 0 {
			c.Y = 0
			c.State = CurtainDown
			return CurtainEventCovered
		}

	case CurtainDown:
		c.Y += c.Speed * dt

		if c.Y >= 60 {
			c.Y = 60
			c.State = CurtainIdle
			return CurtainEventComplete
		}
	}

	return CurtainEventNone
}

func (c *Curtain) Draw(part *engine.FramePart) {
	if c.State == CurtainIdle {
		return
	}

	startY := int(c.Y)
	for y := startY; y < part.Height; y++ {
		for x := 0; x < part.Width; x++ {
			part.SetPixel(x, y, true)
		}
	}
}
