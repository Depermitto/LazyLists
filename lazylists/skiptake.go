package lazylists

func Take[T any](seq Seq[T], n uint) Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		for t := range seq {
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

func (seq SeqIndexed[T]) Take(n uint) SeqIndexed[T] {
	return func(yield func(int, T) bool) {
		i := 0
		for index, t := range seq {
			if uint(i) < n {
				yield(index, t)
			}
			i++
		}
	}
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

func (seq SeqIndexed[T]) TakeWhile(f func(index int, t T) bool) SeqIndexed[T] {
	return func(yield func(int, T) bool) {
		stop := false
		for index, t := range seq {
			if !f(index, t) {
				stop = true
			}

			if !stop {
				yield(index, t)
			}
		}
	}
}

func Skip[T any](seq Seq[T], n uint) Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		for t := range seq {
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

func (seq SeqIndexed[T]) Skip(n uint) SeqIndexed[T] {
	return func(yield func(int, T) bool) {
		i := 0
		for index, t := range seq {
			if uint(i) >= n {
				yield(index, t)
			}
			i++
		}
	}
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

func (seq SeqIndexed[T]) SkipWhile(f func(index int, t T) bool) SeqIndexed[T] {
	return func(yield func(int, T) bool) {
		pass := false
		for index, t := range seq {
			if !f(index, t) {
				pass = true
			}

			if pass {
				yield(index, t)
			}
		}
	}
}
