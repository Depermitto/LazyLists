package lazy

import "iter"

type ZipPair[T, U any] struct {
	First  T
	Second U
}

func zip[T, U any](seq1 Seq[T], seq2 Seq[U]) Seq[ZipPair[T, U]] {
	next1, stop1 := iter.Pull[T](iter.Seq[T](seq1))
	next2, stop2 := iter.Pull[U](iter.Seq[U](seq2))

	return func(yield func(ZipPair[T, U]) bool) {
		for {
			t, ok1 := next1()
			u, ok2 := next2()
			if !ok1 || !ok2 {
				goto stop
			}

			yield(ZipPair[T, U]{First: t, Second: u})
		}
	stop:
		stop1()
		stop2()
	}
}

func Zip[T, U any](sliceOrSeq1 any, sliceOrSeq2 any) Seq[ZipPair[T, U]] {
	return zip[T, U](parseToSeq[T](sliceOrSeq1), parseToSeq[U](sliceOrSeq2))
}
