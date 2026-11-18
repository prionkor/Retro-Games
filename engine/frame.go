package engine

type Frame struct {
	Width      int
	Height     int
	Pixels     []bool
	FrameParts []*FramePart
}

func (f *Frame) MergeFrameParts() {
	// f.Clear()
	for _, part := range f.FrameParts {
		for y := 0; y < part.Height; y++ {
			for x := 0; x < part.Width; x++ {
				frameX := part.X + x
				frameY := part.Y + y

				if frameX < 0 || frameX >= f.Width ||
					frameY < 0 || frameY >= f.Height {
					continue
				}

				partIndex := y*part.Width + x
				frameIndex := frameY*f.Width + frameX

				f.Pixels[frameIndex] =
					f.Pixels[frameIndex] || part.Pixels[partIndex]
			}
		}
	}
}

func (f *Frame) Clear() {
	clear(f.Pixels)
	for _, part := range f.FrameParts {
		part.Clear()
	}
}
