package lazylists

func (seq Seq[T]) Collect() []T {
	s := make([]T, 0)
	for t := range seq {
		s = append(s, t)
	}
	return s
}
