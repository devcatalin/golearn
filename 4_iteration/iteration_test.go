package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		assertCorrectOutput(t, Repeat("a", 5), "aaaaa")
	})

	t.Run("return original repeat times is 0", func(t *testing.T) {
		assertCorrectOutput(t, Repeat("a", 0), "a")
	})

	t.Run("return original repeat times is less than 0", func(t *testing.T) {
		assertCorrectOutput(t, Repeat("a", -5), "a")
	})
}

func assertCorrectOutput(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
