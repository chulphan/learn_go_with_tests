package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("2 plus 2 is 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected %d but got %d", sum, expected)
		}
	})

	t.Run("2 plus 3 is 5", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5

		if sum != expected {
			t.Errorf("expected %d but got %d", sum, expected)
		}
	})
}
