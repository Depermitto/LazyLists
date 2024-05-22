package lazy

import "iter"

type Seq[T any] iter.Seq[T]

func makeSeq[T any](s []T) Seq[T] {
	return func(yield func(T) bool) {
		for _, t := range s {
			yield(t)
		}
	}
}

func parseToSeq[T any](sliceOrSeq any) Seq[T] {
	seq, ok := sliceOrSeq.(Seq[T])
	if ok {
		return seq
	}
	slice, ok := sliceOrSeq.([]T)
	if ok {
		return makeSeq(slice)
	}
	return nil
}
