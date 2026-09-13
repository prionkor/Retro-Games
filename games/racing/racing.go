package racing

type Game struct {
	Player Player
	Track  Track
	Score  int
}

func NewGame() *Game {
	return &Game{
		Player: Player{
			X:      0,
			Y:      0,
			Width:  50,
			Height: 100,
		},
		Track: Track{
			Left:  -100,
			Right: 100,
			Speed: 200,
		},
		Score: 0,
	}
}

type Player struct {
	X float64
	Y float64

	Width  int
	Height int
}

type Track struct {
	Left  float64
	Right float64
	Speed float64
}
