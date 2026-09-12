package engine

import (
	"fmt"

	"github.com/prionkor/retro-games/platform"
)

type Engine struct {
	platform  platform.Platform
	running   bool
	lastFrame float64
	hasFrame  bool
}

func NewEngine(platform platform.Platform) *Engine {
	return &Engine{
		platform: platform,
		running:  false,
		hasFrame: false,
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
	fmt.Println("dt:", dt)
}

func (e *Engine) Render() {
}
