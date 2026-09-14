package browser

import (
	"syscall/js"

	"github.com/prionkor/retro-games/platform"
)

type BrowserPlatform struct {
	canvas   js.Value
	context  js.Value
	input    platform.Input
	keyState map[string]bool
}

var supportedKeys = map[string]bool{
	"ArrowLeft":  true,
	"ArrowRight": true,
}

func NewBrowserPlatform() *BrowserPlatform {
	canvas := js.Global().
		Get("document").
		Call("getElementById", "game")

	context := canvas.Call("getContext", "2d")
	keyState := make(map[string]bool)

	

	b := &BrowserPlatform{
		canvas:   canvas,
		context:  context,
		keyState: keyState,
	}

	b.RegisterKeyboard()

	return b
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
	pixels, width, height = expandPixels(width, height, pixels)

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

func expandPixels(width int, height int, pixels []bool) (
	[]bool,
	int,
	int,
) {
	physicalWidth := width*6 + (width - 1) + 2
	physicalHeight := height*6 + (height - 1) + 2

	expandedPixels := make([]bool, physicalWidth*physicalHeight)

	pattern := []bool{
		true, true, true, true, true, true,
		true, false, false, false, false, true,
		true, false, true, true, false, true,
		true, false, true, true, false, true,
		true, false, false, false, false, true,
		true, true, true, true, true, true,
	}

	for i, pixel := range pixels {
		if !pixel {
			continue
		}

		x := i % width
		y := i / width

		cellX := 1 + x*7
		cellY := 1 + y*7

		for cy := 0; cy < 6; cy++ {
			for cx := 0; cx < 6; cx++ {
				patternIndex := cy*6 + cx

				if !pattern[patternIndex] {
					continue
				}

				index := (cellY+cy)*physicalWidth + (cellX + cx)

				expandedPixels[index] = true
			}
		}
	}

	return expandedPixels, physicalWidth, physicalHeight
}

func (b *BrowserPlatform) IsKeyDown(key string) bool {
	return b.keyState[key]
}

func (b *BrowserPlatform) Input() platform.Input {
	return platform.Input{
		Left:  b.IsKeyDown("ArrowLeft"),
		Right: b.IsKeyDown("ArrowRight"),
	}
}

func (b *BrowserPlatform) RegisterKeyboard() {
	document := js.Global().Get("document")

	keyDown := js.FuncOf(func(this js.Value, args []js.Value) any {
		event := args[0]
		key := event.Get("key").String()

		if !supportedKeys[key] {
			return nil
		}

		b.keyState[key] = true
		return nil
	})

	keyUp := js.FuncOf(func(this js.Value, args []js.Value) any {
		event := args[0]
		key := event.Get("key").String()

		if !supportedKeys[key] {
			return nil
		}

		b.keyState[key] = false
		return nil
	})

	document.Call("addEventListener", "keydown", keyDown)
	document.Call("addEventListener", "keyup", keyUp)
}

var _ platform.Platform = (*BrowserPlatform)(nil)
