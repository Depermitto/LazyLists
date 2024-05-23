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

func (seq SeqIndexed[T]) Filter(f func(int, T) bool) SeqIndexed[T] {
	return func(yield func(int, T) bool) {
		for i, t := range seq {
			if f(i, t) {
				yield(i, t)
			}
		}
	}
}

func (seq Seq[T]) FilterNot(f func(T) bool) Seq[T] {
	return Filter(seq, func(t T) bool { return !f(t) })
}

func (seq SeqIndexed[T]) FilterNot(f func(int, T) bool) SeqIndexed[T] {
	return seq.Filter(func(i int, t T) bool { return !f(i, t) })
}
