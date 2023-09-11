package exercises_test

import (
	exercises "mastering-go-exercises/chapter-1"
	"testing"
)

func TestInvalidWich(t *testing.T) {
	t.Parallel()

	err := exercises.Which()
	if err == nil {
		t.Fatal("want error for empty args")
	}
}
