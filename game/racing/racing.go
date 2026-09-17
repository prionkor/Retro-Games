package racing

import (
	"math/rand"

	"github.com/prionkor/retro-games/engine"
	"github.com/prionkor/retro-games/game"
	"github.com/prionkor/retro-games/platform"
)

type Game struct {
	Car        Car
	Track      Track
	Score      int
	Enemies    []*Car
	MaxEnemies int
	gameOver   bool
	Curtain    game.Curtain
}

func NewGame() *Game {
	return &Game{
		Car:        *NewCar(11, 56, false),
		Track:      *NewTrack(0, 0, 60, 25),
		Enemies:    []*Car{},
		Score:      0,
		MaxEnemies: 5,
		gameOver:   false,
		Curtain:    *game.NewCurtain(40),
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
		g.Car.Pixels(),
	)

	for _, enemy := range g.Enemies {
		f.Draw(
			int(enemy.X),
			int(enemy.Y),
			enemy.Width,
			enemy.Height,
			enemy.Pixels(),
		)
	}

	if !g.gameOver && g.checkCollision() {
		g.EndGame()
		return
	}

	g.removeOldEnemies()
	g.addNewEnemy()

	g.Curtain.Draw(f)
}

func (g *Game) Update(dt float64, input platform.Input) {
	g.Track.Update(dt)

	for _, enemy := range g.Enemies {
		enemy.MoveY(g.Track.Speed * dt)
	}

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

	// curtain on game over

	event := g.Curtain.Animate(dt)

	switch event {
	case game.CurtainEventCovered:
		g.Reset()

	case game.CurtainEventComplete:
		g.Track.Speed = 3
	}
}

func (g *Game) Reset() {
	g.Car.X = 10
	g.Car.Y = 56
	g.Car.Stop = false

	g.Enemies = []*Car{}

	g.Track.Offset = 0
	g.Track.Speed = 0

	g.gameOver = false
}

func (g *Game) removeOldEnemies() {
	remaining := g.Enemies[:0]
	for _, enemy := range g.Enemies {
		if enemy.Y < float64(g.Track.height) {
			remaining = append(remaining, enemy)
		}
	}

	g.Enemies = remaining
}

func (g *Game) addNewEnemy() {
	if len(g.Enemies) >= g.MaxEnemies {
		return
	}

	x := float64(rand.Intn(21) + 1)
	y := float64(-rand.Intn(g.Track.Height()))

	car := NewCar(x, y, true)
	for _, enemy := range g.Enemies {
		if carsOverlap(car, enemy) {
			return
		}
	}
	g.Enemies = append(g.Enemies, car)
}

func carsOverlap(a, b *Car) bool {
	return int(a.X) < int(b.X)+b.Width &&
		int(a.X)+a.Width > int(b.X) &&
		int(a.Y) < int(b.Y)+b.Height &&
		int(a.Y)+a.Height > int(b.Y)
}

func (g *Game) checkCollision() bool {
	for _, enemy := range g.Enemies {
		if carsOverlap(&g.Car, enemy) {
			return true
		}
	}

	return false
}

func (g *Game) EndGame() {
	g.gameOver = true
	g.Track.Speed = 0
	g.Car.Stop = true
	g.Curtain.Trigger()
}
