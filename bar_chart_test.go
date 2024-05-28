package chart

import (
	"bytes"
	"math"
	"testing"
)

func TestBarChartRender(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{
		Width: 1024,
		Title: "Test Title",
		Bars: []Value{
			{Value: 1.0, Label: "One"},
			{Value: 2.0, Label: "Two"},
			{Value: 3.0, Label: "Three"},
			{Value: 4.0, Label: "Four"},
			{Value: 5.0, Label: "Five"},
		},
	}

	buf := bytes.NewBuffer([]byte{})
	err := bc.Render(PNG, buf)
	gotest.AssertNil(t, err)
	gotest.AssertNotZero(t, buf.Len())
}

func TestBarChartRenderZero(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{
		Width: 1024,
		Title: "Test Title",
		Bars: []Value{
			{Value: 0.0, Label: "One"},
			{Value: 0.0, Label: "Two"},
		},
	}

	buf := bytes.NewBuffer([]byte{})
	err := bc.Render(PNG, buf)
	gotest.AssertNotNil(t, err)
}

func TestBarChartProps(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}

	gotest.AssertEqual(t, DefaultDPI, bc.GetDPI())
	bc.DPI = 100
	gotest.AssertEqual(t, 100, bc.GetDPI())

	gotest.AssertNil(t, bc.GetFont())
	f, err := GetDefaultFont()
	gotest.AssertNil(t, err)
	bc.Font = f
	gotest.AssertNotNil(t, bc.GetFont())

	gotest.AssertEqual(t, DefaultChartWidth, bc.GetWidth())
	bc.Width = DefaultChartWidth - 1
	gotest.AssertEqual(t, DefaultChartWidth-1, bc.GetWidth())

	gotest.AssertEqual(t, DefaultChartHeight, bc.GetHeight())
	bc.Height = DefaultChartHeight - 1
	gotest.AssertEqual(t, DefaultChartHeight-1, bc.GetHeight())

	gotest.AssertEqual(t, DefaultBarSpacing, bc.GetBarSpacing())
	bc.BarSpacing = 150
	gotest.AssertEqual(t, 150, bc.GetBarSpacing())

	gotest.AssertEqual(t, DefaultBarWidth, bc.GetBarWidth())
	bc.BarWidth = 75
	gotest.AssertEqual(t, 75, bc.GetBarWidth())
}

func TestBarChartRenderNoBars(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}
	err := bc.Render(PNG, bytes.NewBuffer([]byte{}))
	gotest.AssertNotNil(t, err)
}

func TestBarChartGetRanges(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}

	yr := bc.getRanges()
	gotest.AssertNotNil(t, yr)
	gotest.AssertFalse(t, yr.IsZero())

	gotest.AssertEqual(t, -math.MaxFloat64, yr.GetMax())
	gotest.AssertEqual(t, math.MaxFloat64, yr.GetMin())
}

func TestBarChartGetRangesBarsMinMax(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{
		Bars: []Value{
			{Value: 1.0},
			{Value: 10.0},
		},
	}

	yr := bc.getRanges()
	gotest.AssertNotNil(t, yr)
	gotest.AssertFalse(t, yr.IsZero())

	gotest.AssertEqual(t, 10, yr.GetMax())
	gotest.AssertEqual(t, 1, yr.GetMin())
}

func TestBarChartGetRangesMinMax(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{
		YAxis: YAxis{
			Range: &ContinuousRange{
				Min: 5.0,
				Max: 15.0,
			},
			Ticks: []Tick{
				{Value: 7.0, Label: "Foo"},
				{Value: 11.0, Label: "Foo2"},
			},
		},
		Bars: []Value{
			{Value: 1.0},
			{Value: 10.0},
		},
	}

	yr := bc.getRanges()
	gotest.AssertNotNil(t, yr)
	gotest.AssertFalse(t, yr.IsZero())

	gotest.AssertEqual(t, 15, yr.GetMax())
	gotest.AssertEqual(t, 5, yr.GetMin())
}

func TestBarChartGetRangesTicksMinMax(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{
		YAxis: YAxis{
			Ticks: []Tick{
				{Value: 7.0, Label: "Foo"},
				{Value: 11.0, Label: "Foo2"},
			},
		},
		Bars: []Value{
			{Value: 1.0},
			{Value: 10.0},
		},
	}

	yr := bc.getRanges()
	gotest.AssertNotNil(t, yr)
	gotest.AssertFalse(t, yr.IsZero())

	gotest.AssertEqual(t, 11, yr.GetMax())
	gotest.AssertEqual(t, 7, yr.GetMin())
}

func TestBarChartHasAxes(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}
	gotest.AssertTrue(t, bc.hasAxes())
	bc.YAxis = YAxis{
		Style: Hidden(),
	}
	gotest.AssertFalse(t, bc.hasAxes())
}

func TestBarChartGetDefaultCanvasBox(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}
	b := bc.getDefaultCanvasBox()
	gotest.AssertFalse(t, b.IsZero())
}

func TestBarChartSetRangeDomains(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}
	cb := bc.box()
	yr := bc.getRanges()
	yr2 := bc.setRangeDomains(cb, yr)
	gotest.AssertNotZero(t, yr2.GetDomain())
}

func TestBarChartGetValueFormatters(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{}
	vf := bc.getValueFormatters()
	gotest.AssertNotNil(t, vf)
	gotest.AssertEqual(t, "1234.00", vf(1234.0))

	bc.YAxis.ValueFormatter = func(_ interface{}) string { return "test" }
	gotest.AssertEqual(t, "test", bc.getValueFormatters()(1234))
}

func TestBarChartGetAxesTicks(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{
		Bars: []Value{
			{Value: 1.0},
			{Value: 2.0},
			{Value: 3.0},
		},
	}

	r, err := PNG(128, 128)
	gotest.AssertNil(t, err)
	yr := bc.getRanges()
	yf := bc.getValueFormatters()

	bc.YAxis.Style.Hidden = true
	ticks := bc.getAxesTicks(r, yr, yf)
	gotest.AssertEmpty(t, ticks)

	bc.YAxis.Style.Hidden = false
	ticks = bc.getAxesTicks(r, yr, yf)
	gotest.AssertLen(t, ticks, 2)
}

func TestBarChartCalculateEffectiveBarSpacing(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{
		Width:    1024,
		BarWidth: 10,
		Bars: []Value{
			{Value: 1.0, Label: "One"},
			{Value: 2.0, Label: "Two"},
			{Value: 3.0, Label: "Three"},
			{Value: 4.0, Label: "Four"},
			{Value: 5.0, Label: "Five"},
		},
	}

	spacing := bc.calculateEffectiveBarSpacing(bc.box())
	gotest.AssertNotZero(t, spacing)

	bc.BarWidth = 250
	spacing = bc.calculateEffectiveBarSpacing(bc.box())
	gotest.AssertZero(t, spacing)
}

func TestBarChartCalculateEffectiveBarWidth(t *testing.T) {
	// replaced new assertions helper

	bc := BarChart{
		Width:    1024,
		BarWidth: 10,
		Bars: []Value{
			{Value: 1.0, Label: "One"},
			{Value: 2.0, Label: "Two"},
			{Value: 3.0, Label: "Three"},
			{Value: 4.0, Label: "Four"},
			{Value: 5.0, Label: "Five"},
		},
	}

	cb := bc.box()

	spacing := bc.calculateEffectiveBarSpacing(bc.box())
	gotest.AssertNotZero(t, spacing)

	barWidth := bc.calculateEffectiveBarWidth(bc.box(), spacing)
	gotest.AssertEqual(t, 10, barWidth)

	bc.BarWidth = 250
	spacing = bc.calculateEffectiveBarSpacing(bc.box())
	gotest.AssertZero(t, spacing)
	barWidth = bc.calculateEffectiveBarWidth(bc.box(), spacing)
	gotest.AssertEqual(t, 199, barWidth)

	gotest.AssertEqual(t, cb.Width()+1, bc.calculateTotalBarWidth(barWidth, spacing))

	bw, bs, total := bc.calculateScaledTotalWidth(cb)
	gotest.AssertEqual(t, spacing, bs)
	gotest.AssertEqual(t, barWidth, bw)
	gotest.AssertEqual(t, cb.Width()+1, total)
}

func TestBarChatGetTitleFontSize(t *testing.T) {
	// replaced new assertions helper
	size := BarChart{Width: 2049, Height: 2049}.getTitleFontSize()
	gotest.AssertEqual(t, 48, size)
	size = BarChart{Width: 1025, Height: 1025}.getTitleFontSize()
	gotest.AssertEqual(t, 24, size)
	size = BarChart{Width: 513, Height: 513}.getTitleFontSize()
	gotest.AssertEqual(t, 18, size)
	size = BarChart{Width: 257, Height: 257}.getTitleFontSize()
	gotest.AssertEqual(t, 12, size)
	size = BarChart{Width: 128, Height: 128}.getTitleFontSize()
	gotest.AssertEqual(t, 10, size)
}
