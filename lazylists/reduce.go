package lazylists

func Fold[T, V any](seq Seq[T], init V, f func(acc V, t T) V) V {
	for t := range seq {
		init = f(init, t)
	}
	return init
}

func (seq Seq[T]) Fold(init T, f func(acc T, t T) T) T {
	return Fold(seq, init, f)
}

func Reduce[T any](seq Seq[T], f func(acc T, t T) T) (reduced T, empty bool) {
	empty = true
	for t := range seq {
		if empty {
			empty = false
			reduced = t
		} else {
			reduced = f(reduced, t)
		}
	}
	return
}

func (seq Seq[T]) Reduce(f func(acc T, t T) T) (reduced T, empty bool) {
	return Reduce(seq, f)
}
