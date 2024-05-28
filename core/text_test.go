package chart

import (
	"testing"
)

func TestTextWrapWord(t *testing.T) {
	// replaced new assertions helper

	r, err := PNG(1024, 1024)
	gotest.AssertNil(t, err)
	f, err := GetDefaultFont()
	gotest.AssertNil(t, err)

	basicTextStyle := Style{Font: f, FontSize: 24}

	output := Text.WrapFitWord(r, "this is a test string", 100, basicTextStyle)
	gotest.AssertNotEmpty(t, output)
	gotest.AssertLen(t, output, 3)

	for _, line := range output {
		basicTextStyle.WriteToRenderer(r)
		lineBox := r.MeasureText(line)
		gotest.AssertTrue(t, lineBox.Width() < 100, line+": "+lineBox.String())
	}
	gotest.AssertEqual(t, "this is", output[0])
	gotest.AssertEqual(t, "a test", output[1])
	gotest.AssertEqual(t, "string", output[2])

	output = Text.WrapFitWord(r, "foo", 100, basicTextStyle)
	gotest.AssertLen(t, output, 1)
	gotest.AssertEqual(t, "foo", output[0])

	// test that it handles newlines.
	output = Text.WrapFitWord(r, "this\nis\na\ntest\nstring", 100, basicTextStyle)
	gotest.AssertLen(t, output, 5)

	// test that it handles newlines and long lines.
	output = Text.WrapFitWord(r, "this\nis\na\ntest\nstring that is very long", 100, basicTextStyle)
	gotest.AssertLen(t, output, 8)
}

func TestTextWrapRune(t *testing.T) {
	// replaced new assertions helper

	r, err := PNG(1024, 1024)
	gotest.AssertNil(t, err)
	f, err := GetDefaultFont()
	gotest.AssertNil(t, err)

	basicTextStyle := Style{Font: f, FontSize: 24}

	output := Text.WrapFitRune(r, "this is a test string", 150, basicTextStyle)
	gotest.AssertNotEmpty(t, output)
	gotest.AssertLen(t, output, 2)
	gotest.AssertEqual(t, "this is a t", output[0])
	gotest.AssertEqual(t, "est string", output[1])
}
