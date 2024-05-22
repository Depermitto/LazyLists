package lazy

func take[T any](iter Seq[T], n uint) Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		for t := range iter {
			if uint(i) < n {
				yield(t)
			}
			i++
		}
	}
}

func Take[T any](sliceOrSeq any, n uint) Seq[T] {
	return take(parseToSeq[T](sliceOrSeq), n)
}

func (iter Seq[T]) Take(n uint) Seq[T] {
	return take(iter, n)
}

func skip[T any](iter Seq[T], n uint) Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		for t := range iter {
			if uint(i) >= n {
				yield(t)
			}
			i++
		}
	}
}

func Skip[T any](sliceOrSeq any, n uint) Seq[T] {
	return skip(parseToSeq[T](sliceOrSeq), n)
}

func (iter Seq[T]) Skip(n uint) Seq[T] {
	return skip(iter, n)
}
