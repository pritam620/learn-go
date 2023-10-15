package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("Expected: %q but got %q", expected, repeated)
	}
}

// To run the benchmark: go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
