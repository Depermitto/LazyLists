package lazylists

import "github.com/barweiss/go-tuple"

func (seq Seq[T]) Collect() []T {
	s := make([]T, 0)
	for t := range seq {
		s = append(s, t)
	}
	return s
}

func (seq SeqIndexed[T]) Collect() []tuple.T2[int, T] {
	s := make([]tuple.T2[int, T], 0)
	for i, t := range seq {
		s = append(s, tuple.New2(i, t))
	}
	return s
}
