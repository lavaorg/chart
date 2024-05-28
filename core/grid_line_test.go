package chart

import (
	"testing"
)

func TestGenerateGridLines(t *testing.T) {
	// replaced new assertions helper

	ticks := []Tick{
		{Value: 1.0, Label: "1.0"},
		{Value: 2.0, Label: "2.0"},
		{Value: 3.0, Label: "3.0"},
		{Value: 4.0, Label: "4.0"},
	}

	gl := GenerateGridLines(ticks, Style{}, Style{})
	gotest.AssertLen(t, gl, 2)

	gotest.AssertEqual(t, 2.0, gl[0].Value)
	gotest.AssertEqual(t, 3.0, gl[1].Value)
}
