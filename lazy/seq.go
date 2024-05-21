package lazy

import "iter"

type Seq[T any] iter.Seq[T]

func makeIter[T any](s []T) Seq[T] {
	return func(yield func(T) bool) {
		for _, t := range s {
			yield(t)
		}
	}
}
