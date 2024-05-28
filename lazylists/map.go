package lazylists

import "github.com/barweiss/go-tuple"

func Map[T, V any](seq Seq[T], f func(T) V) Seq[V] {
	return func(yield func(V) bool) {
		for t := range seq {
			yield(f(t))
		}
	}
}

func (seq Seq[T]) Map(f func(T) T) Seq[T] {
	return Map(seq, f)
}

func (seq Seq2[I, T]) Map(f func(I, T) tuple.T2[I, T]) Seq2[I, T] {
	return func(yield func(I, T) bool) {
		for i, t := range seq {
			both := f(i, t)
			yield(both.V1, both.V2)
		}
	}
}
