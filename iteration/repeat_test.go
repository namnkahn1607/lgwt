package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("numbered times of repeat", func(t *testing.T) {
		repeated := Repeat("a", 6)
		expected := "aaaaaa"
		assertCorrectRepeated(t, repeated, expected)
	})

	t.Run("negative times of repeat", func(t *testing.T) {
		repeated := Repeat("a", -2)
		expected := "aaaaa"
		assertCorrectRepeated(t, repeated, expected)
	})
}

func assertCorrectRepeated(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("expected %s, got %s", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output:
	// aaaaa
}
