package racing

type Track struct {
	X      int
	Y      int
	height int
	width  int

	Pattern       []bool
	PatternHeight int
	Pixels        []bool

	Speed  float64
	Offset float64
}

func NewTrack(x, y, height, width int) *Track {
	patternHeight := 3
	pattern := make([]bool, width*patternHeight)
	pixels := make([]bool, width*height)

	for row := range patternHeight {
		for col := range width {
			if row < 2 && (col < 1 || col >= width-1) {
				pattern[row*width+col] = true
			}
		}
	}

	return &Track{
		X:             x,
		Y:             y,
		height:        height,
		width:         width,
		Pattern:       pattern,
		PatternHeight: patternHeight,
		Pixels:        pixels,
		Speed:         9,
		Offset:        0,
	}
}

func (t *Track) Update(dt float64) {
	t.Offset += t.Speed * dt

	for y := 0; y < t.height; y++ {
		patternY := (y - int(t.Offset)) % t.PatternHeight

		if patternY < 0 {
			patternY += t.PatternHeight
		}

		for x := 0; x < t.width; x++ {
			patternIndex := patternY*t.width + x
			pixelIndex := y*t.width + x

			t.Pixels[pixelIndex] = t.Pattern[patternIndex]
		}
	}
}

func (t *Track) Height() int {
	return t.height
}

func (t *Track) Width() int {
	return t.width
}
