package lazylists

import (
	"slices"
	"testing"
)

const testSliceSize = 200

func TestTakeWhile(t *testing.T) {
	s := Range(0, testSliceSize).Map(func(int) int { return 4 }).Collect()
	s[testSliceSize/2] = -10

	iter := TakeWhile(MakeSeq(s), func(i int) bool { return i != -10 })
	got := iter.Collect()

	expected := s[:10]
	if !slices.Equal(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func BenchmarkTakeWhile(b *testing.B) {
	s := Range(0, testSliceSize).Map(func(int) int { return 4 }).Collect()
	s[testSliceSize/2] = -10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for t := range TakeWhile(s, func(i int) bool { return i != -10 }) {
			_ = t
		}
	}
}
