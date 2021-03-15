package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times a if given empty string", func(t *testing.T) {
		repeated := Repeat("", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("got %q but expect %q", repeated, expected)
		}
	})

	t.Run("repeat 5 times given string", func(t *testing.T) {
		repeated := Repeat("c", 5)
		expected := "ccccc"

		if repeated != expected {
			t.Errorf("got %q but expect %q", repeated, expected)
		}
	})

	t.Run("반복횟수를 인자로 전달하여 전달된 인자만큼 문자열을 반복하도록", func(t *testing.T) {
		repeated := Repeat("c", 10)
		expected := "cccccccccc"

		if repeated != expected {
			t.Errorf("got %q but expect %q", repeated, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
