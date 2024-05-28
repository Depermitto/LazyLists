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

func (seq Seq2[I, T]) Filter(f func(I, T) bool) Seq2[I, T] {
	return func(yield func(I, T) bool) {
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

func (seq Seq2[I, T]) FilterNot(f func(I, T) bool) Seq2[I, T] {
	return seq.Filter(func(i I, t T) bool { return !f(i, t) })
}
