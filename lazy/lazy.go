package lazy

import (
	"Depermitto/LazyLists/lazylists"
	"github.com/barweiss/go-tuple"
)

func parseToSeq[T any](sliceOrSeq any) lazylists.Seq[T] {
	seq, ok := sliceOrSeq.(lazylists.Seq[T])
	if ok {
		return seq
	}
	slice, ok := sliceOrSeq.([]T)
	if ok {
		return lazylists.MakeSeq(slice)
	}
	return nil
}

func All[T any](sliceOrSeq any, f func(T) bool) bool {
	return lazylists.All(parseToSeq[T](sliceOrSeq), f)
}

func Any[T any](sliceOrSeq any, f func(T) bool) bool {
	return lazylists.Any(parseToSeq[T](sliceOrSeq), f)
}

func Enumerate[T any](sliceOrSeq any) lazylists.SeqIndexed[T] {
	return lazylists.Enumerate(parseToSeq[T](sliceOrSeq))
}

func Filter[T any](sliceOrSeq any, f func(T) bool) lazylists.Seq[T] {
	return lazylists.Filter(parseToSeq[T](sliceOrSeq), f)
}

func FilterNot[T any](sliceOrSeq any, f func(T) bool) lazylists.Seq[T] {
	return lazylists.Filter(parseToSeq[T](sliceOrSeq), func(t T) bool { return !f(t) })
}

func Map[T, V any](sliceOrSeq any, f func(T) V) lazylists.Seq[V] {
	return lazylists.Map(parseToSeq[T](sliceOrSeq), f)
}

func MinMaxBy[T any](sliceOrSeq any, cmp func(T, T) int) (min T, max T, empty bool) {
	return lazylists.MinMaxBy(parseToSeq[T](sliceOrSeq), cmp)
}

func MinBy[T any](sliceOrSeq any, cmp func(T, T) int) (min T, empty bool) {
	min, _, empty = lazylists.MinMaxBy(parseToSeq[T](sliceOrSeq), cmp)
	return
}

func MaxBy[T any](sliceOrSeq any, cmp func(T, T) int) (max T, empty bool) {
	_, max, empty = lazylists.MinMaxBy(parseToSeq[T](sliceOrSeq), cmp)
	return
}

func Generator[T any](init T, next func(T) (T, error)) lazylists.Seq[T] {
	return lazylists.Generator(init, next)
}

func Range(a int, b int) lazylists.Seq[int] {
	return lazylists.Range(a, b)
}

func StepBy[T any](sliceOrSeq any, step uint) lazylists.Seq[T] {
	return lazylists.StepBy(parseToSeq[T](sliceOrSeq), step)
}

func Fold[T, V any](sliceOrSeq any, init V, f func(acc V, t T) V) V {
	return lazylists.Fold(parseToSeq[T](sliceOrSeq), init, f)
}

func Reduce[T any](sliceOrSeq any, f func(acc T, t T) T) (reduced T, empty bool) {
	return lazylists.Reduce(parseToSeq[T](sliceOrSeq), f)
}

func Rev[T any](s []T) lazylists.Seq[T] {
	return lazylists.Rev(s)
}

func Take[T any](sliceOrSeq any, n uint) lazylists.Seq[T] {
	return lazylists.Take(parseToSeq[T](sliceOrSeq), n)
}

func TakeWhile[T any](sliceOrSeq any, f func(T) bool) lazylists.Seq[T] {
	return lazylists.TakeWhile(parseToSeq[T](sliceOrSeq), f)
}

func Skip[T any](sliceOrSeq any, n uint) lazylists.Seq[T] {
	return lazylists.Skip(parseToSeq[T](sliceOrSeq), n)
}

func SkipWhile[T any](sliceOrSeq any, f func(T) bool) lazylists.Seq[T] {
	return lazylists.SkipWhile(parseToSeq[T](sliceOrSeq), f)
}

func Zip[T, U any](sliceOrSeq1 any, sliceOrSeq2 any) lazylists.Seq[tuple.T2[T, U]] {
	return lazylists.Zip[T, U](parseToSeq[T](sliceOrSeq1), parseToSeq[U](sliceOrSeq2))
}
