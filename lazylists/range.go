package lazylists

func Generator[T any](init T, next func(T) (T, error)) Seq[T] {
	return func(yield func(T) bool) {
		var err error = nil
		for ; err == nil; init, err = next(init) {
			yield(init)
		}
	}
}

func Range(a int, b int) Seq[int] {
	return func(yield func(int) bool) {
		for i := a; i < b; i++ {
			yield(i)
		}
	}
}

func StepBy[T any](seq Seq[T], step uint) Seq[T] {
	if step == 0 {
		return nil
	}

	return func(yield func(T) bool) {
		i := step
		for t := range seq {
			if i == step {
				yield(t)
				i = 1
			} else {
				i++
			}
		}
	}
}

func (seq Seq[T]) StepBy(step uint) Seq[T] {
	return StepBy(seq, step)
}

func (seq Seq2[I, T]) StepBy(step uint) Seq2[I, T] {
	if step == 0 {
		return nil
	}

	return func(yield func(I, T) bool) {
		i := step
		for index, t := range seq {
			if i == step {
				yield(index, t)
				i = 1
			} else {
				i++
			}
		}
	}
}
