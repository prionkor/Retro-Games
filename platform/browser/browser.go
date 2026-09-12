package browser

import (
	"syscall/js"

	"github.com/prionkor/retro-games/platform"
)

type BrowserPlatform struct{}

func (b *BrowserPlatform) Log(message string) {
	js.Global().Get("console").Call("log", message)
}

func (b *BrowserPlatform) RequestFrame(callback func(timestamp float64)) {
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) any {
		defer cb.Release() // Release the callback after it's called
		timestamp := args[0].Float()
		callback(timestamp)
		return nil
	})

	js.Global().Call("requestAnimationFrame", cb)
}

var _ platform.Platform = (*BrowserPlatform)(nil)
