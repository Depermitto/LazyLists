package lazy

func mapTo[T, V any](seq Seq[T], f func(T) V) Seq[V] {
	return func(yield func(V) bool) {
		for t := range seq {
			yield(f(t))
		}
	}
}

func Map[T, V any](sliceOrSeq any, f func(T) V) Seq[V] {
	return mapTo(parseToSeq[T](sliceOrSeq), f)
}

func (seq Seq[T]) Map(f func(T) T) Seq[T] {
	return mapTo(seq, f)
}
