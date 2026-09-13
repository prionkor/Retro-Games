package engine

type Frame struct {
	Width  int
	Height int
	Pixels []bool
}

func (f *Frame) SetPixel(x int, y int, value bool) {
	if x < 0 || x >= f.Width || y < 0 || y >= f.Height {
		return
	}

	f.Pixels[y*f.Width+x] = value
}

func (f *Frame) FillRect(x, y, width, height int, value bool) {
	for py := range height {
		for px := range width {
			f.SetPixel(x+px, y+py, value)
		}
	}
}

func (f *Frame) Draw(
	x, y int,
	width, height int,
	pixels []bool,
) {
	for py := range height {
		for px := range width {
			frameIndex := (y+py)*f.Width + (x + px)
			pixelIndex := py*width + px

			f.Pixels[frameIndex] =
				f.Pixels[frameIndex] || pixels[pixelIndex]
		}
	}
}

func (f *Frame) Clear() {
	for i := range f.Pixels {
		f.Pixels[i] = false
	}
}
