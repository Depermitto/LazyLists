package lazylists

import (
	"slices"
	"testing"
)

const testSliceSize = 200

func TestTakeWhile(t *testing.T) {
	s := make([]int, testSliceSize)
	for i := 0; i < testSliceSize; i++ {
		s[i] = 4
	}
	s[testSliceSize/2] = -10

	iter := TakeWhile(MakeSeq(s), func(i int) bool { return i != -10 })
	got := iter.Collect()

	expected := s[:testSliceSize/2]
	if !slices.Equal(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func BenchmarkTakeWhile(b *testing.B) {
	s := Range(0, testSliceSize).Map(func(int) int { return 4 }).Collect()
	s[testSliceSize/2] = -10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for t := range TakeWhile(MakeSeq(s), func(i int) bool { return i != -10 }) {
			_ = t
		}
	}
}
