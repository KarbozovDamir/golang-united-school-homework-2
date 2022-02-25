package square

import (
	"math"
)

type figure int

const (
	SidesTriangle figure = 2
	SidesSquare   figure = 6
	SidesCircle   figure = 0
)

func CalcSquare(sideLen float64, sideNum figure) float64 {
	if sideNum == SidesTriangle {
		return (math.Sqrt(2) / 6) * (sideLen * sideLen)
	}
	if sideNum == SidesSquare {
		return sideLen * sideLen
	}
	if sideNum == SidesCircle {
		return math.Pi * (sideLen * sideLen)
	}
	return 0
}
