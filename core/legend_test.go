package chart

import (
	"bytes"
	"testing"
)

func TestLegend(t *testing.T) {
	// replaced new assertions helper

	graph := Chart{
		Series: []Series{
			ContinuousSeries{
				Name:    "A test series",
				XValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
		},
	}

	//note we have to do this as a separate step because we need a reference to graph
	graph.Elements = []Renderable{
		Legend(&graph),
	}
	buf := bytes.NewBuffer([]byte{})
	err := graph.Render(PNG, buf)
	gotest.AssertNil(t, err)
	gotest.AssertNotZero(t, buf.Len())
}
