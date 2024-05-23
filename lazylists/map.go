package lazylists

func Map[T, V any](seq Seq[T], f func(T) V) Seq[V] {
	return func(yield func(V) bool) {
		for t := range seq {
			yield(f(t))
		}
	}
}

func (seq Seq[T]) Map(f func(T) T) Seq[T] {
	return Map(seq, f)
}
