package lazy

func filter[T any](seq Seq[T], f func(T) bool) Seq[T] {
	return func(yield func(T) bool) {
		for t := range seq {
			if f(t) {
				yield(t)
			}
		}
	}
}

func Filter[T any](sliceOrSeq any, f func(T) bool) Seq[T] {
	return filter(parseToSeq[T](sliceOrSeq), f)
}

func (seq Seq[T]) Filter(f func(T) bool) Seq[T] {
	return filter(seq, f)
}

func FilterNot[T any](sliceOrSeq any, f func(T) bool) Seq[T] {
	return filter(parseToSeq[T](sliceOrSeq), func(t T) bool { return !f(t) })
}

func (seq Seq[T]) FilterNot(f func(T) bool) Seq[T] {
	return filter(seq, func(t T) bool { return !f(t) })
}
