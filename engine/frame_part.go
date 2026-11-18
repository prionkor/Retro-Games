package engine

type FramePart struct {
	X          int // X is the coordinate that frame part should start in main frame
	Y          int // Y is the coordinate that frame part should start in main frame
	Width      int
	Height     int
	Pixels     []bool
	PixelBlock PixelBlock
}

func (f *FramePart) SetPixel(x int, y int, value bool) {
	if x < 0 || x >= f.Width || y < 0 || y >= f.Height {
		return
	}

	f.Pixels[y*f.Width+x] = value
}

func (f *FramePart) FillRect(x, y, width, height int, value bool) {
	for py := range height {
		for px := range width {
			f.SetPixel(x+px, y+py, value)
		}
	}
}

func (f *FramePart) Draw(
	x, y int,
	width, height int,
	pixels []bool,
) {
	for py := range height {
		frameY := y + py
		if frameY < 0 || frameY >= f.Height {
			continue
		}

		for px := range width {
			frameX := x + px

			if frameX < 0 || frameX >= f.Width {
				continue
			}

			frameIndex := frameY*f.Width + frameX
			pixelIndex := py*width + px

			f.Pixels[frameIndex] =
				f.Pixels[frameIndex] || pixels[pixelIndex]
		}
	}
}

func (f *FramePart) Clear() {
	for i := range f.Pixels {
		f.Pixels[i] = false
	}
}
