package vulture

type List[T any] []T

func FromString(val string) List[byte] {
	bytes := []byte(val)
	return FromSlice(bytes)
}

func FromSlice[T any](val []T) List[T] {
	var s List[T] = val
	return s
}

func reverse[T any](list *List[T]) {
	ll := *list
	for i := len(ll)/2 - 1; i >= 0; i-- {
		opp := len(ll) - 1 - i
		ll[i], ll[opp] = ll[opp], ll[i]
	}
}

func (s List[T]) IntoString() string {
	var ss string

	switch t := any(s).(type) {
	case List[byte]:
		ss = string(t)
	}

	return ss
}

func (s List[T]) Iter() iterator[T] {
	return iterator[T]{s}
}

// func (s *List[T]) IterRef() iterator[T] {
// 	return iterator[T]{s}
// }
