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

func (seq SeqIndexed[T]) Map(f func(int, T) tuple.T2[int, T]) SeqIndexed[T] {
	return func(yield func(int, T) bool) {
		for i, t := range seq {
			both := f(i, t)
			yield(both.V1, both.V2)
		}
	}
}
