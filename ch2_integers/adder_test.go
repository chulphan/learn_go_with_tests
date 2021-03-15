package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {

	assertEquals := func(t testing.TB, sum, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected %d but got %d", sum, expected)
		}
	}

	t.Run("2 plus 2 is 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertEquals(t, sum, expected)
	})

	t.Run("2 plus 3 is 5", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5

		assertEquals(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(3, 5)
	fmt.Println(sum)
	// Output: 8
}
