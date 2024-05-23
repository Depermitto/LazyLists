package lazylists

func MinMaxBy[T any](seq Seq[T], cmp func(T, T) int) (min T, max T, empty bool) {
	empty = true
	for t := range seq {
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

func (seq Seq[T]) MinMaxBy(cmp func(T, T) int) (min T, max T, empty bool) {
	return MinMaxBy(seq, cmp)
}

func (seq Seq[T]) MinBy(cmp func(T, T) int) (min T, empty bool) {
	min, _, empty = MinMaxBy(seq, cmp)
	return
}

func (seq Seq[T]) MaxBy(cmp func(T, T) int) (max T, empty bool) {
	_, max, empty = MinMaxBy(seq, cmp)
	return
}
