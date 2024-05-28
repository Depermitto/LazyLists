package lazylists

import "iter"

type Seq[T any] iter.Seq[T]
type Seq2[I, T any] iter.Seq2[I, T]

func MakeSeq[T any](s []T) Seq[T] {
	return func(yield func(T) bool) {
		for _, t := range s {
			yield(t)
		}
	}
}
