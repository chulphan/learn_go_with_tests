package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat()
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("got %q but expect %q", repeated, expected)
	}
}
