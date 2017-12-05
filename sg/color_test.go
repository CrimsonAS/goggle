package sg

import "testing"

func assertColorString(t *testing.T, index int, c Color, s string) {
	cs := c.String()
	if cs != s {
		t.Fatalf("For color %d, expected string %s, got string %s", index, s, cs)
	}
}

type colorTest struct {
	color          Color
	expectedString string
}

func TestColorString(t *testing.T) {
	tests := []colorTest{
		colorTest{color: Color{1, 0, 0, 0}, expectedString: "#ff000000"},
		colorTest{color: Color{0, 1, 0, 0}, expectedString: "#00ff0000"},
		colorTest{color: Color{0, 0, 1, 0}, expectedString: "#0000ff00"},
		colorTest{color: Color{0, 0, 0, 1}, expectedString: "#000000ff"},

		colorTest{color: Color{0.5, 0, 0, 0}, expectedString: "#7f000000"},
		colorTest{color: Color{0, 0.5, 0, 0}, expectedString: "#007f0000"},
		colorTest{color: Color{0, 0, 0.5, 0}, expectedString: "#00007f00"},
		colorTest{color: Color{0, 0, 0, 0.5}, expectedString: "#0000007f"},

		colorTest{color: Color{0.1, 0, 0, 0}, expectedString: "#19000000"},
		colorTest{color: Color{0, 0.1, 0, 0}, expectedString: "#00190000"},
		colorTest{color: Color{0, 0, 0.1, 0}, expectedString: "#00001900"},
		colorTest{color: Color{0, 0, 0, 0.1}, expectedString: "#00000019"},

		colorTest{color: Color{1, 1, 1, 1}, expectedString: "#ffffffff"},
		colorTest{color: Color{0, 0, 0, 0}, expectedString: "#00000000"},
	}

	for idx, test := range tests {
		assertColorString(t, idx, test.color, test.expectedString)
	}
}
