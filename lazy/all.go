package lazy

func allMatch[T any](seq Seq[T], f func(T) bool) bool {
	for t := range seq {
		if !f(t) {
			return false
		}
	}
	return true
}

func All[T any](sliceOrSeq any, f func(T) bool) bool {
	return allMatch(parseToSeq[T](sliceOrSeq), f)
}

func (seq Seq[T]) All(f func(T) bool) bool {
	return allMatch(seq, f)
}

func anyMatch[T any](seq Seq[T], f func(T) bool) bool {
	// can it be !allMatch(seq, !f)?
	for t := range seq {
		if f(t) {
			return true
		}
	}
	return false
}

func Any[T any](sliceOrSeq any, f func(T) bool) bool {
	return anyMatch(parseToSeq[T](sliceOrSeq), f)
}

func (seq Seq[T]) Any(f func(T) bool) bool {
	return anyMatch(seq, f)
}
