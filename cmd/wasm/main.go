//go:build js && wasm

package main

import (
	"github.com/prionkor/retro-games/engine"
	"github.com/prionkor/retro-games/games/racing"
	"github.com/prionkor/retro-games/platform/browser"
)

func main() {
	p := browser.NewBrowserPlatform()
	g := racing.NewGame()
	e := engine.NewEngine(p, g)

	e.Start()

	select {}
}
