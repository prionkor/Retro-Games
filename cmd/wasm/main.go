package main

import (
	"github.com/prionkor/retro-games/engine"
	"github.com/prionkor/retro-games/platform/browser"
)

func main() {
	p := &browser.BrowserPlatform{}
	e := engine.NewEngine(p)
	p.Log("Hello from Go WebAssembly!")

	e.Start()

	select {}
}
