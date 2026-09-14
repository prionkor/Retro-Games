package platform

type Platform interface {
	Log(message string)
	RequestFrame(callback func(timestamp float64))
	Present(width int, height int, pixels []bool)

	IsKeyDown(key string) bool
	Input() Input
}

type Input struct {
	Left  bool
	Right bool
}
