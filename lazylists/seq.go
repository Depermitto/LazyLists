package lazylists

import "iter"

type Seq[T any] iter.Seq[T]
type SeqIndexed[T any] iter.Seq2[int, T]

func MakeSeq[T any](s []T) Seq[T] {
	return func(yield func(T) bool) {
		for _, t := range s {
			yield(t)
		}
	}
}
