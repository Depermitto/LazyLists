package lazy

func mapTo[T, V any](iter Seq[T], f func(T) V) Seq[V] {
	return func(yield func(V) bool) {
		for t := range iter {
			yield(f(t))
		}
	}
}

func Map[T, V any](s []T, f func(T) V) Seq[V] {
	return mapTo(makeSeq(s), f)
}

func (iter Seq[T]) Map(f func(T) T) Seq[T] {
	return mapTo(iter, f)
}
