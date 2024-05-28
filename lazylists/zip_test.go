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
		for i1, i2 := range Zip[int, int](MakeSeq(s1), MakeSeq(s2)) {
			_ = i1
			_ = i2
		}
	}
}
