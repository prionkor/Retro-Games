package racing

type Car struct {
	X      float64
	Y      float64
	Width  int
	Height int
	pixels []bool
	FlipY  bool
	Stop   bool
}

func NewCar(x, y float64, flipY bool) *Car {
	return &Car{
		X:      x,
		Y:      y,
		Width:  3,
		Height: 4,
		FlipY:  flipY,
		pixels: []bool{
			false, true, false,
			true, true, true,
			false, true, false,
			true, false, true,
		},
		Stop: false,
	}
}

func (c *Car) MoveX(amount float64) {
	if c.Stop {
		return
	}
	c.X += amount
}

func (c *Car) MoveY(amount float64) {
	if c.Stop {
		return
	}
	c.Y += amount
}

func (c *Car) Pixels() []bool {
	if !c.FlipY {
		return c.pixels
	}

	rotated := make([]bool, len(c.pixels))

	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			sourceX := c.Width - 1 - x
			sourceY := c.Height - 1 - y

			sourceIndex := sourceY*c.Width + sourceX
			targetIndex := y*c.Width + x

			rotated[targetIndex] = c.pixels[sourceIndex]
		}
	}

	return rotated
}
