package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times a if given empty string", func(t *testing.T) {
		repeated := Repeat("")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("got %q but expect %q", repeated, expected)
		}
	})

	t.Run("repeat 5 times given string", func(t *testing.T) {
		repeated := Repeat("c")
		expected := "ccccc"

		if repeated != expected {
			t.Errorf("got %q but expect %q", repeated, expected)
		}
	})
}
