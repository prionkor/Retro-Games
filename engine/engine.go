package engine

import (
	"github.com/prionkor/retro-games/platform"
)

type Engine struct {
	platform platform.Platform
	Frame    Frame
	game     Game

	running   bool
	lastFrame float64
	hasFrame  bool
}

func NewEngine(platform platform.Platform, game Game) *Engine {
	return &Engine{
		platform: platform,
		running:  false,
		hasFrame: false,
		game:     game,

		Frame: Frame{
			Width:  40,
			Height: 60,
			Pixels: make([]bool, 40*60),
		},
	}
}

func (e *Engine) Start() {
	e.running = true
	e.platform.RequestFrame(e.frame)
}

func (e *Engine) frame(timestamp float64) {
	if !e.hasFrame {
		e.lastFrame = timestamp
		e.hasFrame = true
	}

	dt := (timestamp - e.lastFrame) / 1000
	e.lastFrame = timestamp

	e.Update(dt)
	e.Render()

	if e.running {
		e.platform.RequestFrame(e.frame)
	}
}

func (e *Engine) Update(dt float64) {
	input := e.platform.Input()
	e.game.Update(dt, input)
}

func (e *Engine) Render() {
	e.game.Render(&e.Frame)
	e.platform.Present(e.Frame.Width, e.Frame.Height, e.Frame.Pixels)
}
