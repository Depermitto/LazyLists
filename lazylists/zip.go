package lazylists

import (
	"github.com/barweiss/go-tuple"
	"iter"
)

func Zip[T, U any](seq1 Seq[T], seq2 Seq[U]) Seq[tuple.T2[T, U]] {
	next1, stop1 := iter.Pull[T](iter.Seq[T](seq1))
	next2, stop2 := iter.Pull[U](iter.Seq[U](seq2))

	return func(yield func(tuple.T2[T, U]) bool) {
		for {
			t, ok1 := next1()
			u, ok2 := next2()
			if !ok1 || !ok2 {
				goto stop
			}

			yield(tuple.New2(t, u))
		}
	stop:
		stop1()
		stop2()
	}
}
