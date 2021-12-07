package numbers

import "testing"

func TestFormatSuffix(t *testing.T) {
	cases := []struct {
		Title  string
		Input  int64
		Output string
	}{
		{
			Title:  "with a Billion",
			Input:  400300200100,
			Output: "400.30B",
		},
		{
			Title:  "with a million",
			Input:  300200100,
			Output: "300.20M",
		},
		{
			Title:  "with a thousand",
			Input:  200100,
			Output: "200.10K",
		},
	}

	for _, test := range cases {
		got := FormatSuffix(test.Input)
		want := test.Output

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}
