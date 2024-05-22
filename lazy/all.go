package lazy

func allMatch[T any](iter Seq[T], f func(T) bool) bool {
	for t := range iter {
		if !f(t) {
			return false
		}
	}
	return true
}

func All[T any](sliceOrSeq any, f func(T) bool) bool {
	return allMatch(parseToSeq[T](sliceOrSeq), f)
}

func (iter Seq[T]) All(f func(T) bool) bool {
	return allMatch(iter, f)
}

func anyMatch[T any](iter Seq[T], f func(T) bool) bool {
	// can it be !allMatch(iter, !f)?
	for t := range iter {
		if f(t) {
			return true
		}
	}
	return false
}

func Any[T any](sliceOrSeq any, f func(T) bool) bool {
	return anyMatch(parseToSeq[T](sliceOrSeq), f)
}

func (iter Seq[T]) Any(f func(T) bool) bool {
	return anyMatch(iter, f)
}
