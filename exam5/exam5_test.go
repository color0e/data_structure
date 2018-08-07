package exam5

import (
	"testing"
)

func TestFib(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{0, 0},
		{5, 5},
		{6, 7},
	}
	for _, c := range cases {
		got := Fib(c.in)
		if got != c.want {
			t.Errorf("Fib(%d) == %d, want %d", c.in, got, c.want)
		}
	}
}
