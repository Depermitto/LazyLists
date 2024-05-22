package lazy

import (
	"math/rand/v2"
	"slices"
	"testing"
)

const testSliceSize = 200

func TestTakeWhile(t *testing.T) {
	s := rand.Perm(testSliceSize)
	slices.Sort(s)
	s[10] = -10

	iter := TakeWhile(makeSeq(s), func(i int) bool { return i != -10 })
	got := iter.Collect()

	expected := s[:10]
	if !slices.Equal(got, expected) {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}

func BenchmarkTakeWhile(b *testing.B) {
	s := Range(0, testSliceSize).Map(func(int) int { return 4 }).Collect()
	s[0] = -10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for t := range TakeWhile(s, func(i int) bool { return i != -10 }) {
			_ = t
		}
	}
}

func BenchmarkTakeWhileOld(b *testing.B) {
	s := Range(0, testSliceSize).Map(func(int) int { return 4 }).Collect()
	s[0] = -10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for t := range TakeWhileOld(s, func(i int) bool { return i != -10 }) {
			_ = t
		}
	}
}
