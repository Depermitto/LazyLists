package lazylists

func All[T any](seq Seq[T], f func(T) bool) bool {
	for t := range seq {
		if !f(t) {
			return false
		}
	}
	return true
}

func (seq Seq[T]) All(f func(T) bool) bool {
	return All(seq, f)
}

func (seq Seq2[T, U]) All(f func(T, U) bool) bool {
	for i, t := range seq {
		if !f(i, t) {
			return false
		}
	}
	return true
}

func Any[T any](seq Seq[T], f func(T) bool) bool {
	// can it be !All(seq, !f)?
	for t := range seq {
		if f(t) {
			return true
		}
	}
	return false
}

func (seq Seq[T]) Any(f func(T) bool) bool {
	return Any(seq, f)
}

func (seq Seq2[I, T]) Any(f func(I, T) bool) bool {
	for i, t := range seq {
		if f(i, t) {
			return true
		}
	}
	return false
}
