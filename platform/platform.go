package platform

type Platform interface {
	Log(message string)
	RequestFrame(callback func(timestamp float64))
}
