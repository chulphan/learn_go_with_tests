package arrays_and_slices

import "testing"

func TestSum(t *testing.T) {

	assertEqual := func(t testing.TB, got, want int, numbers []int) {
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	}

	t.Run("collection of 5 number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertEqual(t, got, want, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		assertEqual(t, got, want, numbers)
	})
}
