package drawing

import (
	"fmt"
	"testing"

	"image/color"
)

func TestColorFromHex(t *testing.T) {
	white := ColorFromHex("FFFFFF")
	gotest.AssertEqual(t, ColorWhite, white)

	shortWhite := ColorFromHex("FFF")
	gotest.AssertEqual(t, ColorWhite, shortWhite)

	black := ColorFromHex("000000")
	gotest.AssertEqual(t, ColorBlack, black)

	shortBlack := ColorFromHex("000")
	gotest.AssertEqual(t, ColorBlack, shortBlack)

	red := ColorFromHex("FF0000")
	gotest.AssertEqual(t, ColorRed, red)

	shortRed := ColorFromHex("F00")
	gotest.AssertEqual(t, ColorRed, shortRed)

	green := ColorFromHex("008000")
	gotest.AssertEqual(t, ColorGreen, green)

	// shortGreen := ColorFromHex("0F0")
	// gotest.AssertEqual(t, ColorGreen, shortGreen)

	blue := ColorFromHex("0000FF")
	gotest.AssertEqual(t, ColorBlue, blue)

	shortBlue := ColorFromHex("00F")
	gotest.AssertEqual(t, ColorBlue, shortBlue)
}

func TestColorFromHex_handlesHash(t *testing.T) {
	withHash := ColorFromHex("#FF0000")
	gotest.AssertEqual(t, ColorRed, withHash)

	withoutHash := ColorFromHex("#FF0000")
	gotest.AssertEqual(t, ColorRed, withoutHash)
}

func TestColorFromAlphaMixedRGBA(t *testing.T) {
	black := ColorFromAlphaMixedRGBA(color.Black.RGBA())
	gotest.AssertTrue(t, black.Equals(ColorBlack), black.String())

	white := ColorFromAlphaMixedRGBA(color.White.RGBA())
	gotest.AssertTrue(t, white.Equals(ColorWhite), white.String())
}

func Test_ColorFromRGBA(t *testing.T) {
	value := "rgba(192, 192, 192, 1.0)"
	parsed := ColorFromRGBA(value)
	gotest.AssertEqual(t, ColorSilver, parsed)

	value = "rgba(192,192,192,1.0)"
	parsed = ColorFromRGBA(value)
	gotest.AssertEqual(t, ColorSilver, parsed)

	value = "rgba(192,192,192,1.5)"
	parsed = ColorFromRGBA(value)
	gotest.AssertEqual(t, ColorSilver, parsed)
}

func TestParseColor(t *testing.T) {
	testCases := [...]struct {
		Input    string
		Expected Color
	}{
		{"", Color{}},
		{"white", ColorWhite},
		{"WHITE", ColorWhite}, // caps!
		{"black", ColorBlack},
		{"red", ColorRed},
		{"green", ColorGreen},
		{"blue", ColorBlue},
		{"silver", ColorSilver},
		{"maroon", ColorMaroon},
		{"purple", ColorPurple},
		{"fuchsia", ColorFuchsia},
		{"lime", ColorLime},
		{"olive", ColorOlive},
		{"yellow", ColorYellow},
		{"navy", ColorNavy},
		{"teal", ColorTeal},
		{"aqua", ColorAqua},

		{"rgba(192, 192, 192, 1.0)", ColorSilver},
		{"rgba(192,192,192,1.0)", ColorSilver},
		{"rgb(192, 192, 192)", ColorSilver},
		{"rgb(192,192,192)", ColorSilver},

		{"#FF0000", ColorRed},
		{"#008000", ColorGreen},
		{"#0000FF", ColorBlue},
		{"#F00", ColorRed},
		{"#080", Color{0, 136, 0, 255}},
		{"#00F", ColorBlue},
	}

	for index, tc := range testCases {
		actual := ParseColor(tc.Input)
		gotest.AssertEqual(t, tc.Expected, actual, fmt.Sprintf("test case: %d -> %s", index, tc.Input))
	}
}
