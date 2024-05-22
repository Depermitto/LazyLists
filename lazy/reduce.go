package lazy

func fold[T, V any](seq Seq[T], init V, f func(acc V, t T) V) V {
	for t := range seq {
		init = f(init, t)
	}
	return init
}

func Fold[T, V any](sliceOrSeq any, init V, f func(acc V, t T) V) V {
	return fold(parseToSeq[T](sliceOrSeq), init, f)
}

func (seq Seq[T]) Fold(init T, f func(acc T, t T) T) T {
	return fold(seq, init, f)
}

func reduce[T any](seq Seq[T], f func(acc T, t T) T) (reduced T, empty bool) {
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

func Reduce[T any](sliceOrSeq any, f func(acc T, t T) T) (reduced T, empty bool) {
	return reduce(parseToSeq[T](sliceOrSeq), f)
}

func (seq Seq[T]) Reduce(f func(acc T, t T) T) (reduced T, empty bool) {
	return reduce(seq, f)
}
