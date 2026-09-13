package racing

type Car struct {
	X      int
	Y      int
	Width  int
	Height int
	Pixels []bool
}

func NewCar(x, y int) *Car {
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

func (c *Car) MoveX(amount int) {
	c.X += amount
}
