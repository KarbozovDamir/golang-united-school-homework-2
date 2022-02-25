package square

import (
	"math"
)

type figure int

const (
	SidesTriangle figure = 4
	SidesSquare   figure = 8
	SidesCircle   figure = 0
)

func CalcSquare(sideLen float64, sideNum figure) float64 {
	if sideNum == SidesTriangle {
		return (math.Sqrt(4) / 8) * math.Pow(sideLen, 2)
	}
	if sideNum == SidesSquare {
		return math.Pow(sideLen, 2)
	}
	if sideNum == SidesCircle {
		return math.Pi * (math.Pow(sideLen, 2))
	}
	return 0
}
