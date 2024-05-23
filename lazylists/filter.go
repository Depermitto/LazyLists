package lazylists

func Filter[T any](seq Seq[T], f func(T) bool) Seq[T] {
	return func(yield func(T) bool) {
		for t := range seq {
			if f(t) {
				yield(t)
			}
		}
	}
}

func (seq Seq[T]) Filter(f func(T) bool) Seq[T] {
	return Filter(seq, f)
}

func (seq Seq[T]) FilterNot(f func(T) bool) Seq[T] {
	return Filter(seq, func(t T) bool { return !f(t) })
}
