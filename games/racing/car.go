package racing

type Car struct {
	X      float64
	Y      float64
	Width  int
	Height int
	Pixels []bool
}

func NewCar(x, y float64) *Car {
	return &Car{
		X:      x,
		Y:      y,
		Width:  3,
		Height: 4,
		Pixels: []bool{
			false, true, false,
			true, true, true,
			false, true, false,
			true, false, true,
		},
	}
}

func (c *Car) MoveX(amount float64) {
	c.X += amount
}
