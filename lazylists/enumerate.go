package lazylists

import "github.com/barweiss/go-tuple"

type Indexed[T any] tuple.T2[int, T]

func Enumerate[T any](seq Seq[T]) SeqIndexed[T] {
	return func(yield func(int, T) bool) {
		i := 0
		for t := range seq {
			yield(i, t)
			i++
		}
	}
}

func (seq Seq[T]) Enumerate() SeqIndexed[T] {
	return Enumerate(seq)
}

func (seq SeqIndexed[T]) UnIndex() Seq[T] {
	return func(yield func(T) bool) {
		for _, t := range seq {
			yield(t)
		}
	}
}
