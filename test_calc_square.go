package square

import (
	"testing"
)
var (
	sides FigureSide
	sideLength float64

)

func TestCalcSquareFunction(t *testing.T) {
	sides = 4
	sideLength = 1.0
	sq := CalcSquare(sideLength, sides)
	if sq != 1.0 {
		t.Error("Incorrect square!")
	}
}

