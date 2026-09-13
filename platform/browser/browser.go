package browser

import (
	"syscall/js"

	"github.com/prionkor/retro-games/platform"
)

type BrowserPlatform struct {
	canvas  js.Value
	context js.Value
}

func NewBrowserPlatform() *BrowserPlatform {
	canvas := js.Global().
		Get("document").
		Call("getElementById", "game")

	context := canvas.Call("getContext", "2d")

	return &BrowserPlatform{
		canvas:  canvas,
		context: context,
	}
}

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
func (b *BrowserPlatform) Present(width int, height int, pixels []bool) {
	imageData := b.context.Call("createImageData", width, height)
	data := imageData.Get("data")

	for i := 0; i < len(pixels); i++ {
		if pixels[i] {
			data.SetIndex(i*4, 0)
			data.SetIndex(i*4+1, 0)
			data.SetIndex(i*4+2, 0)
			data.SetIndex(i*4+3, 255)
		} else {
			data.SetIndex(i*4, 187)
			data.SetIndex(i*4+1, 187)
			data.SetIndex(i*4+2, 187)
			data.SetIndex(i*4+3, 255)
		}
	}

	b.context.Call("putImageData", imageData, 0, 0)
}

func (b *BrowserPlatform) IsKeyDown(key string) bool {
	keyState := js.Global().Get("keyState")
	if keyState.IsUndefined() {
		return false
	}

	return keyState.Get(key).Bool()
}

var _ platform.Platform = (*BrowserPlatform)(nil)
