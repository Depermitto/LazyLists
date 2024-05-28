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

func (seq Seq2[I, T]) Take(n uint) Seq2[I, T] {
	return func(yield func(I, T) bool) {
		j := 0
		for i, t := range seq {
			if uint(j) < n {
				yield(i, t)
			}
			j++
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

func (seq Seq2[I, T]) TakeWhile(f func(I, T) bool) Seq2[I, T] {
	return func(yield func(I, T) bool) {
		stop := false
		for i, t := range seq {
			if !f(i, t) {
				stop = true
			}

			if !stop {
				yield(i, t)
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

func (seq Seq2[I, T]) Skip(n uint) Seq2[I, T] {
	return func(yield func(I, T) bool) {
		j := 0
		for i, t := range seq {
			if uint(j) >= n {
				yield(i, t)
			}
			j++
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

func (seq Seq2[I, T]) SkipWhile(f func(I, T) bool) Seq2[I, T] {
	return func(yield func(I, T) bool) {
		pass := false
		for i, t := range seq {
			if !f(i, t) {
				pass = true
			}

			if pass {
				yield(i, t)
			}
		}
	}
}
