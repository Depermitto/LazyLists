package lazy

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

func stepBy[T any](seq Seq[T], step uint) Seq[T] {
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

func StepBy[T any](sliceOrSeq any, step uint) Seq[T] {
	return stepBy(parseToSeq[T](sliceOrSeq), step)
}

func (seq Seq[T]) StepBy(step uint) Seq[T] {
	return stepBy(seq, step)
}
