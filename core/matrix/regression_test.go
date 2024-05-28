package matrix

import (
	"testing"
)

func TestPoly(t *testing.T) {
	// replaced new assertions helper
	var xGiven = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var yGiven = []float64{1, 6, 17, 34, 57, 86, 121, 162, 209, 262, 321}
	var degree = 2

	c, err := Poly(xGiven, yGiven, degree)
	gotest.AssertNil(t, err)
	gotest.AssertLen(t, c, 3)

	gotest.AssertInDelta(t, c[0], 0.999999999, DefaultEpsilon)
	gotest.AssertInDelta(t, c[1], 2, DefaultEpsilon)
	gotest.AssertInDelta(t, c[2], 3, DefaultEpsilon)
}
