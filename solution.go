package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"math"
)

type FigureSide int

const (
	SidesCircle FigureSide = 0
	SidesTriangle FigureSide = 3
	SidesSquare FigureSide = 4
)

func CalcSquare(sideLen float64, sidesNum FigureSide) float64 {
	
	if sidesNum == SidesCircle {
		return math.Pi * sideLen * sideLen
	} 

	if sidesNum == SidesTriangle {
		return sideLen * sideLen * math.Sqrt(3) / 4
	} 

	if sidesNum == SidesSquare {
		return sideLen * sideLen
	} 

	return 0
}
