package lazylists

import "github.com/barweiss/go-tuple"

func (seq Seq[T]) Collect() []T {
	s := make([]T, 0)
	for t := range seq {
		s = append(s, t)
	}
	return s
}

func (seq Seq2[I, T]) Collect() []tuple.T2[I, T] {
	s := make([]tuple.T2[I, T], 0)
	for i, t := range seq {
		s = append(s, tuple.New2(i, t))
	}
	return s
}
