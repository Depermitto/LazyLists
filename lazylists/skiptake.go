package lazylists

func Take[T any](iter Seq[T], n uint) Seq[T] {
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

func (seq Seq[T]) Take(n uint) Seq[T] {
	return Take(seq, n)
}

func TakeWhile[T any](seq Seq[T], f func(T) bool) Seq[T] {
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

func (seq Seq[T]) TakeWhile(f func(T) bool) Seq[T] {
	return TakeWhile(seq, f)
}

func Skip[T any](iter Seq[T], n uint) Seq[T] {
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

func (seq Seq[T]) Skip(n uint) Seq[T] {
	return Skip(seq, n)
}

func SkipWhile[T any](seq Seq[T], f func(T) bool) Seq[T] {
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

func (seq Seq[T]) SkipWhile(f func(T) bool) Seq[T] {
	return SkipWhile(seq, f)
}
