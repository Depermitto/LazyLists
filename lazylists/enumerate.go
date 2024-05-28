package lazylists

func Enumerate[T any](seq Seq[T]) Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for t := range seq {
			yield(i, t)
			i++
		}
	}
}

func (seq Seq[T]) Enumerate() Seq2[int, T] {
	return Enumerate(seq)
}

func (seq Seq2[I, T]) UnIndex() Seq[T] {
	return func(yield func(T) bool) {
		for _, t := range seq {
			yield(t)
		}
	}
}
