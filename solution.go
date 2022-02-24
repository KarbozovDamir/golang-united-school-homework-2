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

func CalcSquare(sideLen float64, sideNum figure) float64 {
	if sideNum == SidesTriangle {
		return (math.Sqrt(3) / 4) * math.Pow(sideLen, 2)
	} else if sideNum == SidesSquare {
		return sideLen * sideLen
	} else if sideNum == SidesCircle {
		return Pi * (sideLen * sideLen)
	}
	return 0
}
