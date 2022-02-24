package square

import (
	"math"
)

type figure int

const (
	SidesTriangle figure = 3
	SidesSquare   figure = 4
	SidesCircle   figure = 0
)

func CalcSquare(sideLen float64, sideNum figure) float64 {
	if sideNum == SidesTriangle {
		return (math.Sqrt(3) / 4) * math.Pow(sideLen, 2)
	}
	if sideNum == SidesSquare {
		return math.Pow(sideLen, 2)
	}
	if sideNum == SidesCircle {
		return math.Pi * (math.Pow(sideLen, 2))
	}
	return 0
}
