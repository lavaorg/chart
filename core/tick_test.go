package chart

import (
	"testing"
)

func TestGenerateContinuousTicks(t *testing.T) {
	// replaced new assertions helper

	f, err := GetDefaultFont()
	gotest.AssertNil(t, err)

	r, err := PNG(1024, 1024)
	gotest.AssertNil(t, err)
	r.SetFont(f)

	ra := &ContinuousRange{
		Min:    0.0,
		Max:    10.0,
		Domain: 256,
	}

	vf := FloatValueFormatter

	ticks := GenerateContinuousTicks(r, ra, false, Style{}, vf)
	gotest.AssertNotEmpty(t, ticks)
	gotest.AssertLen(t, ticks, 11)
	gotest.AssertEqual(t, 0.0, ticks[0].Value)
	gotest.AssertEqual(t, 10, ticks[len(ticks)-1].Value)
}

func TestGenerateContinuousTicksDescending(t *testing.T) {
	// replaced new assertions helper

	f, err := GetDefaultFont()
	gotest.AssertNil(t, err)

	r, err := PNG(1024, 1024)
	gotest.AssertNil(t, err)
	r.SetFont(f)

	ra := &ContinuousRange{
		Min:        0.0,
		Max:        10.0,
		Domain:     256,
		Descending: true,
	}

	vf := FloatValueFormatter

	ticks := GenerateContinuousTicks(r, ra, false, Style{}, vf)
	gotest.AssertNotEmpty(t, ticks)
	gotest.AssertLen(t, ticks, 11)
	gotest.AssertEqual(t, 10.0, ticks[0].Value)
	gotest.AssertEqual(t, 9.0, ticks[1].Value)
	gotest.AssertEqual(t, 1.0, ticks[len(ticks)-2].Value)
	gotest.AssertEqual(t, 0.0, ticks[len(ticks)-1].Value)
}
