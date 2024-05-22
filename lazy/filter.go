package lazy

func filter[T any](iter Seq[T], f func(T) bool) Seq[T] {
	return func(yield func(T) bool) {
		for t := range iter {
			if f(t) {
				yield(t)
			}
		}
	}
}

func Filter[T any](sliceOrSeq any, f func(T) bool) Seq[T] {
	return filter(parseToSeq[T](sliceOrSeq), f)
}

func (iter Seq[T]) Filter(f func(T) bool) Seq[T] {
	return filter(iter, f)
}

func (iter Seq[T]) FilterNot(f func(T) bool) Seq[T] {
	return filter(iter, func(t T) bool { return !f(t) })
}
