package lazy

func minMaxBy[T any](iter Seq[T], cmp func(T, T) int) (min T, max T, empty bool) {
	empty = true
	for t := range iter {
		if empty {
			empty = false
			min = t
			max = t
		} else if cmp(t, min) < 0 {
			min = t
		} else if cmp(t, max) > 0 {
			max = t
		}
	}
	return
}

func MinMaxBy[T any](s []T, cmp func(T, T) int) (min T, max T, empty bool) {
	return minMaxBy(makeSeq(s), cmp)
}

func MinBy[T any](s []T, cmp func(T, T) int) (min T, empty bool) {
	min, _, empty = minMaxBy(makeSeq(s), cmp)
	return
}

func MaxBy[T any](s []T, cmp func(T, T) int) (max T, empty bool) {
	_, max, empty = minMaxBy(makeSeq(s), cmp)
	return
}

func (iter Seq[T]) MinMaxBy(cmp func(T, T) int) (min T, max T, empty bool) {
	return minMaxBy(iter, cmp)
}

func (iter Seq[T]) MinBy(cmp func(T, T) int) (min T, empty bool) {
	min, _, empty = minMaxBy(iter, cmp)
	return
}

func (iter Seq[T]) MaxBy(cmp func(T, T) int) (max T, empty bool) {
	_, max, empty = minMaxBy(iter, cmp)
	return
}
