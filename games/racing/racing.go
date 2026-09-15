package racing

import (
	"github.com/prionkor/retro-games/engine"
	"github.com/prionkor/retro-games/platform"
)

type Game struct {
	Car   Car
	Track Track
	Score int
}

func NewGame() *Game {
	return &Game{
		Car:   *NewCar(11, 56),
		Track: *NewTrack(0, 0, 60, 25),
		Score: 0,
	}
}

func (g *Game) Render(f *engine.Frame) {
	f.Clear()

	f.Draw(
		g.Track.X,
		g.Track.Y,
		g.Track.Width(),
		g.Track.Height(),
		g.Track.Pixels,
	)

	f.Draw(
		int(g.Car.X),
		int(g.Car.Y),
		g.Car.Width,
		g.Car.Height,
		g.Car.Pixels,
	)
}

func (g *Game) Update(dt float64, input platform.Input) {
	g.Track.Update(dt)

	minX := g.Track.X + 1
	maxX := g.Track.X + g.Track.Width() - 1 - g.Car.Width

	if input.Left {
		g.Car.MoveX(-1)

		if g.Car.X < float64(minX) {
			g.Car.X = float64(minX)
		}
	}

	if input.Right {
		g.Car.MoveX(1)

		if g.Car.X > float64(maxX) {
			g.Car.X = float64(maxX)
		}
	}
}
