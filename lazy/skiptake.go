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

func (seq Seq[T]) Take(n uint) Seq[T] {
	return take(seq, n)
}

func takeWhile[T any](seq Seq[T], f func(T) bool) Seq[T] {
	return func(yield func(T) bool) {
		stop := false
		for t := range seq {
			if !f(t) {
				stop = true
			}

			if !stop {
				yield(t)
			}
		}
	}
}

func TakeWhile[T any](sliceOrSeq any, f func(T) bool) Seq[T] {
	return takeWhile(parseToSeq[T](sliceOrSeq), f)
}

func (seq Seq[T]) TakeWhile(f func(T) bool) Seq[T] {
	return takeWhile(seq, f)
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

func (seq Seq[T]) Skip(n uint) Seq[T] {
	return skip(seq, n)
}

func skipWhile[T any](seq Seq[T], f func(T) bool) Seq[T] {
	return func(yield func(T) bool) {
		pass := false
		for t := range seq {
			if !f(t) {
				pass = true
			}

			if pass {
				yield(t)
			}
		}
	}
}

func SkipWhile[T any](sliceOrSeq any, f func(T) bool) Seq[T] {
	return skipWhile(parseToSeq[T](sliceOrSeq), f)
}

func (seq Seq[T]) SkipWhile(f func(T) bool) Seq[T] {
	return skipWhile(seq, f)
}
