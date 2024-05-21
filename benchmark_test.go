package main

import (
	"Depermitto/LazyLists/lazy"
	"math/rand/v2"
	"mtoohey.com/iter/v2"
	"testing"
)

const testSliceSize = 100

func BenchmarkFilterEvenIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := rand.Perm(testSliceSize)
		even := iter.Elems(s).Filter(func(i int) bool {
			return i%2 == 0
		})
		for range even.Collect() {
			// LOOP BODY
		}
	}
}

func BenchmarkFilterEvenLazyLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := rand.Perm(testSliceSize)
		even := lazy.Filter(s, func(i int) bool {
			return i%2 == 0
		})
		for _ = range even {
			// LOOP BODY
		}
	}
}
