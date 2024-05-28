package lazylists

import "github.com/barweiss/go-tuple"

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

func minMaxBy[I, T any](seq Seq2[I, T], cmp func(I, T, I, T) int) (min tuple.T2[I, T], max tuple.T2[I, T], empty bool) {
	empty = true
	for i, t := range seq {
		if empty {
			empty = false
			min = tuple.New2(i, t)
			max = tuple.New2(i, t)
		} else if cmp(i, t, min.V1, min.V2) < 0 {
			min = tuple.New2(i, t)
		} else if cmp(i, t, max.V1, max.V2) > 0 {
			max = tuple.New2(i, t)
		}
	}
	return
}

func (seq Seq2[I, T]) MinMaxBy(cmp func(I, T, I, T) int) (min tuple.T2[I, T], max tuple.T2[I, T], empty bool) {
	return minMaxBy(seq, cmp)
}

func (seq Seq2[I, T]) MinBy(cmp func(I, T, I, T) int) (min tuple.T2[I, T], empty bool) {
	min, _, empty = minMaxBy(seq, cmp)
	return
}

func (seq Seq2[I, T]) MaxBy(cmp func(I, T, I, T) int) (max tuple.T2[I, T], empty bool) {
	_, max, empty = minMaxBy(seq, cmp)
	return
}
