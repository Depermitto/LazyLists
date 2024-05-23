package lazylists

import (
	"math/rand/v2"
	"testing"
)

var (
	s1 = rand.Perm(testSliceSize)
	s2 = rand.Perm(testSliceSize)
)

func BenchmarkZip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for pair := range Zip[int, int](s1, s2) {
			_ = pair
		}
	}
}
