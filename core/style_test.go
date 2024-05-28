package chart

import (
	"testing"

	"github.com/golang/freetype/truetype"
	"github.com/lavaorg/chart/core/drawing"
)

func TestStyleIsZero(t *testing.T) {
	// replaced new assertions helper
	zero := Style{}
	gotest.AssertTrue(t, zero.IsZero())

	strokeColor := Style{StrokeColor: drawing.ColorWhite}
	gotest.AssertFalse(t, strokeColor.IsZero())

	fillColor := Style{FillColor: drawing.ColorWhite}
	gotest.AssertFalse(t, fillColor.IsZero())

	strokeWidth := Style{StrokeWidth: 5.0}
	gotest.AssertFalse(t, strokeWidth.IsZero())

	fontSize := Style{FontSize: 12.0}
	gotest.AssertFalse(t, fontSize.IsZero())

	fontColor := Style{FontColor: drawing.ColorWhite}
	gotest.AssertFalse(t, fontColor.IsZero())

	font := Style{Font: &truetype.Font{}}
	gotest.AssertFalse(t, font.IsZero())
}

func TestStyleGetStrokeColor(t *testing.T) {
	// replaced new assertions helper

	unset := Style{}
	gotest.AssertEqual(t, drawing.ColorTransparent, unset.GetStrokeColor())
	gotest.AssertEqual(t, drawing.ColorWhite, unset.GetStrokeColor(drawing.ColorWhite))

	set := Style{StrokeColor: drawing.ColorWhite}
	gotest.AssertEqual(t, drawing.ColorWhite, set.GetStrokeColor())
	gotest.AssertEqual(t, drawing.ColorWhite, set.GetStrokeColor(drawing.ColorBlack))
}

func TestStyleGetFillColor(t *testing.T) {
	// replaced new assertions helper

	unset := Style{}
	gotest.AssertEqual(t, drawing.ColorTransparent, unset.GetFillColor())
	gotest.AssertEqual(t, drawing.ColorWhite, unset.GetFillColor(drawing.ColorWhite))

	set := Style{FillColor: drawing.ColorWhite}
	gotest.AssertEqual(t, drawing.ColorWhite, set.GetFillColor())
	gotest.AssertEqual(t, drawing.ColorWhite, set.GetFillColor(drawing.ColorBlack))
}

func TestStyleGetStrokeWidth(t *testing.T) {
	// replaced new assertions helper

	unset := Style{}
	gotest.AssertEqual(t, DefaultStrokeWidth, unset.GetStrokeWidth())
	gotest.AssertEqual(t, DefaultStrokeWidth+1, unset.GetStrokeWidth(DefaultStrokeWidth+1))

	set := Style{StrokeWidth: DefaultStrokeWidth + 2}
	gotest.AssertEqual(t, DefaultStrokeWidth+2, set.GetStrokeWidth())
	gotest.AssertEqual(t, DefaultStrokeWidth+2, set.GetStrokeWidth(DefaultStrokeWidth+1))
}

func TestStyleGetFontSize(t *testing.T) {
	// replaced new assertions helper

	unset := Style{}
	gotest.AssertEqual(t, DefaultFontSize, unset.GetFontSize())
	gotest.AssertEqual(t, DefaultFontSize+1, unset.GetFontSize(DefaultFontSize+1))

	set := Style{FontSize: DefaultFontSize + 2}
	gotest.AssertEqual(t, DefaultFontSize+2, set.GetFontSize())
	gotest.AssertEqual(t, DefaultFontSize+2, set.GetFontSize(DefaultFontSize+1))
}

func TestStyleGetFontColor(t *testing.T) {
	// replaced new assertions helper

	unset := Style{}
	gotest.AssertEqual(t, drawing.ColorTransparent, unset.GetFontColor())
	gotest.AssertEqual(t, drawing.ColorWhite, unset.GetFontColor(drawing.ColorWhite))

	set := Style{FontColor: drawing.ColorWhite}
	gotest.AssertEqual(t, drawing.ColorWhite, set.GetFontColor())
	gotest.AssertEqual(t, drawing.ColorWhite, set.GetFontColor(drawing.ColorBlack))
}

func TestStyleGetFont(t *testing.T) {
	// replaced new assertions helper

	f, err := GetDefaultFont()
	gotest.AssertNil(t, err)

	unset := Style{}
	gotest.AssertNil(t, unset.GetFont())
	gotest.AssertEqual(t, f, unset.GetFont(f))

	set := Style{Font: f}
	gotest.AssertNotNil(t, set.GetFont())
}

func TestStyleGetPadding(t *testing.T) {
	// replaced new assertions helper

	unset := Style{}
	gotest.AssertTrue(t, unset.GetPadding().IsZero())
	gotest.AssertFalse(t, unset.GetPadding(DefaultBackgroundPadding).IsZero())
	gotest.AssertEqual(t, DefaultBackgroundPadding, unset.GetPadding(DefaultBackgroundPadding))

	set := Style{Padding: DefaultBackgroundPadding}
	gotest.AssertFalse(t, set.GetPadding().IsZero())
	gotest.AssertEqual(t, DefaultBackgroundPadding, set.GetPadding())
	gotest.AssertEqual(t, DefaultBackgroundPadding, set.GetPadding(Box{
		Top:    DefaultBackgroundPadding.Top + 1,
		Left:   DefaultBackgroundPadding.Left + 1,
		Right:  DefaultBackgroundPadding.Right + 1,
		Bottom: DefaultBackgroundPadding.Bottom + 1,
	}))
}

func TestStyleWithDefaultsFrom(t *testing.T) {
	// replaced new assertions helper

	f, err := GetDefaultFont()
	gotest.AssertNil(t, err)

	unset := Style{}
	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Font:        f,
		Padding:     DefaultBackgroundPadding,
	}

	coalesced := unset.InheritFrom(set)
	gotest.AssertEqual(t, set, coalesced)
}

func TestStyleGetStrokeOptions(t *testing.T) {
	// replaced new assertions helper

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Padding:     DefaultBackgroundPadding,
	}
	svgStroke := set.GetStrokeOptions()
	gotest.AssertFalse(t, svgStroke.StrokeColor.IsZero())
	gotest.AssertNotZero(t, svgStroke.StrokeWidth)
	gotest.AssertTrue(t, svgStroke.FillColor.IsZero())
	gotest.AssertTrue(t, svgStroke.FontColor.IsZero())
}

func TestStyleGetFillOptions(t *testing.T) {
	// replaced new assertions helper

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Padding:     DefaultBackgroundPadding,
	}
	svgFill := set.GetFillOptions()
	gotest.AssertFalse(t, svgFill.FillColor.IsZero())
	gotest.AssertZero(t, svgFill.StrokeWidth)
	gotest.AssertTrue(t, svgFill.StrokeColor.IsZero())
	gotest.AssertTrue(t, svgFill.FontColor.IsZero())
}

func TestStyleGetFillAndStrokeOptions(t *testing.T) {
	// replaced new assertions helper

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Padding:     DefaultBackgroundPadding,
	}
	svgFillAndStroke := set.GetFillAndStrokeOptions()
	gotest.AssertFalse(t, svgFillAndStroke.FillColor.IsZero())
	gotest.AssertNotZero(t, svgFillAndStroke.StrokeWidth)
	gotest.AssertFalse(t, svgFillAndStroke.StrokeColor.IsZero())
	gotest.AssertTrue(t, svgFillAndStroke.FontColor.IsZero())
}

func TestStyleGetTextOptions(t *testing.T) {
	// replaced new assertions helper

	set := Style{
		StrokeColor: drawing.ColorWhite,
		StrokeWidth: 5.0,
		FillColor:   drawing.ColorWhite,
		FontColor:   drawing.ColorWhite,
		Padding:     DefaultBackgroundPadding,
	}
	svgStroke := set.GetTextOptions()
	gotest.AssertTrue(t, svgStroke.StrokeColor.IsZero())
	gotest.AssertZero(t, svgStroke.StrokeWidth)
	gotest.AssertTrue(t, svgStroke.FillColor.IsZero())
	gotest.AssertFalse(t, svgStroke.FontColor.IsZero())
}
