package lazy

func (iter Seq[T]) Collect() []T {
	s := make([]T, 0)
	for t := range iter {
		s = append(s, t)
	}
	return s
}
