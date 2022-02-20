package solution

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type myint int

const SidesTriangle myint = 3
const SidesSquare myint = 4
const SidesCircle myint = 0

func CalcSquare(sideLen float64, sidesNum myint) float64 {
	var Pi64 float64 = math.Pi
	var S float64

	if sidesNum == SidesTriangle {
		S = sideLen * 2

	} else if sidesNum == SidesSquare {

		S = sideLen * sideLen

	} else if sidesNum == SidesCircle {

		S = Pi64 * (sideLen * sideLen)

	}

	return S
}
