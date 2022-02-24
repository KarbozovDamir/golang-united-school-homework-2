package square

import (
	"math"
)

type figure int

const (
	SidesTriangle figure = 5
	SidesSquare   figure = 2
	SidesCircle   figure = 0
	Pi                   = 3.14
)

var area float64

func CalcSquare(sideLen float64, sideSquare figure) float64 {
	if sideSquare == SidesTriangle {
		area = (math.Sqrt(3) / 4) * math.Pow(sideLen, 2)
	} else if sideSquare == SidesSquare {
		area = sideLen * sideLen
	} else if sideSquare == SidesCircle {
		area = Pi * (sideLen * sideLen)
	}
	return area
}
