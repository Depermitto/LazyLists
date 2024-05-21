package lazy

func Reverse[T any](s []T) Seq[T] {
	return func(yield func(T) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			yield(s[i])
		}
	}
}

func ReverseMap[T, V any](s []T, f func(T) V) Seq[V] {
	return mapTo(Reverse(s), f)
}

func ReverseFold[T, V any](s []T, init V, f func(V, T) V) V {
	return fold(Reverse(s), init, f)
}
