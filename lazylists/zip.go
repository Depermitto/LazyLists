package lazylists

import (
	"iter"
)

func Zip[I, T any](seq1 Seq[I], seq2 Seq[T]) Seq2[I, T] {
	next1, stop1 := iter.Pull[I](iter.Seq[I](seq1))
	next2, stop2 := iter.Pull[T](iter.Seq[T](seq2))

	return func(yield func(I, T) bool) {
		for {
			i, ok1 := next1()
			t, ok2 := next2()
			if !ok1 || !ok2 {
				goto stop
			}

			yield(i, t)
		}
	stop:
		stop1()
		stop2()
	}
}
