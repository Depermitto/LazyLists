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
