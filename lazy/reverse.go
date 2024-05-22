package lazy

func Rev[T any](s []T) Seq[T] {
	return func(yield func(T) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			yield(s[i])
		}
	}
}
