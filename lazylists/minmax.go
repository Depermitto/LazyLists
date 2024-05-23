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

func minMaxBy[T any](seq SeqIndexed[T], cmp func(int, T, int, T) int) (min Indexed[T], max Indexed[T], empty bool) {
	empty = true
	for i, t := range seq {
		if empty {
			empty = false
			min = Indexed[T]{V1: i, V2: t}
			max = Indexed[T]{V1: i, V2: t}
		} else if cmp(i, t, min.V1, min.V2) < 0 {
			min = Indexed[T]{V1: i, V2: t}
		} else if cmp(i, t, max.V1, max.V2) > 0 {
			max = Indexed[T]{V1: i, V2: t}
		}
	}
	return
}

func (seq SeqIndexed[T]) MinMaxBy(cmp func(int, T, int, T) int) (min Indexed[T], max Indexed[T], empty bool) {
	return minMaxBy(seq, cmp)
}

func (seq SeqIndexed[T]) MinBy(cmp func(int, T, int, T) int) (min Indexed[T], empty bool) {
	min, _, empty = minMaxBy(seq, cmp)
	return
}

func (seq SeqIndexed[T]) MaxBy(cmp func(int, T, int, T) int) (max Indexed[T], empty bool) {
	_, max, empty = minMaxBy(seq, cmp)
	return
}
