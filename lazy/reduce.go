package lazy

func fold[T, V any](iter Seq[T], init V, f func(acc V, t T) V) V {
	for t := range iter {
		init = f(init, t)
	}
	return init
}

func Fold[T, V any](sliceOrSeq any, init V, f func(acc V, t T) V) V {
	return fold(parseToSeq[T](sliceOrSeq), init, f)
}

func (iter Seq[T]) Fold(init T, f func(acc T, t T) T) T {
	return fold(iter, init, f)
}

func reduce[T any](iter Seq[T], f func(acc T, t T) T) (reduced T, empty bool) {
	empty = true
	for t := range iter {
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

func (iter Seq[T]) Reduce(f func(acc T, t T) T) (reduced T, empty bool) {
	return reduce(iter, f)
}
