package vulture

import (
	"golang.org/x/exp/constraints"
)

type ranges[T constraints.Integer] [2]T

func Ranges[T constraints.Integer](start, end T) ranges[T] {
	return ranges[T]{start, end}
}

func (r ranges[T]) Exclusive() ranges[T] {
	if r[0] < r[1] {
		return ranges[T]{r[0] + 1, r[1] - 1}
	}

	return ranges[T]{r[0] - 1, r[1] + 1}
}

func (r ranges[T]) Collect() List[T] {
	return r.build()
}

func (r ranges[T]) build() List[T] {
	if r[0] < r[1] {
		return r.buildIncr()
	}

	return r.buildDecr()
}

func (r ranges[T]) buildIncr() List[T] {
	list := List[T]{}

	for i := r[0]; i <= r[1]; i++ {
		list = append(list, i)
	}

	return list
}

func (r ranges[T]) buildDecr() List[T] {
	list := List[T]{}

	for i := r[0]; i >= r[1]; i-- {
		list = append(list, i)
	}

	return list
}
