package lazylists

func Enumerate[T any](seq Seq[T]) Indexed[T] {
	return func(yield func(int, T) bool) {
		i := 0
		for t := range seq {
			yield(i, t)
			i++
		}
	}
}

func (seq Seq[T]) Enumerate() Indexed[T] {
	return Enumerate(seq)
}
