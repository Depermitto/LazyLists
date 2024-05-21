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

func Take[T any](s []T, n uint) Seq[T] {
	return take(makeIter(s), n)
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

func Skip[T any](s []T, n uint) Seq[T] {
	return skip(makeIter(s), n)
}

func (iter Seq[T]) Skip(n uint) Seq[T] {
	return skip(iter, n)
}
