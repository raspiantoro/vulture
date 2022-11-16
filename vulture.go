package vulture

type numbers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type hofable interface {
	numbers
}

type List[T hofable] []T

func FromString(val string) List[byte] {
	bytes := []byte(val)
	return IntoSlice(bytes)
}

func IntoSlice[T hofable](val []T) List[T] {
	var s List[T] = val
	return s
}

func Min[T hofable](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func Max[T hofable](a, b T) T {
	if a < b {
		return b
	}

	return a
}

func (s List[T]) ScanRight(init T, fn func(accumulator T, val T) T) List[T] {
	var ss List[T]
	var acc T = init

	for i := range s {
		acc = fn(acc, s[len(s)-1-i])
		ss = append(List[T]{acc}, ss...)
	}

	return ss
}

func (s List[T]) ScanLeft(init T, fn func(accumulator, val T) T) List[T] {
	var ss List[T]
	var acc T = init

	for _, v := range s {
		acc = fn(acc, v)
		ss = append(ss, acc)
	}

	return ss
}

func (s List[T]) FoldRight(init T, fn func(accumulator, val T) T) T {
	var result T = init

	for i := range s {
		result = fn(s[len(s)-1-i], result)
	}

	return result
}

func (s List[T]) FoldLeft(init T, fn func(accumulator, val T) T) T {
	var result T = init

	for _, v := range s {
		result = fn(result, v)
	}

	return result
}

func (s List[T]) Filter(fn func(T) bool) List[T] {
	ss := List[T]{}

	for _, v := range s {
		if fn(v) {
			ss = append(ss, v)
		}
	}

	return ss
}

func (s List[T]) IntoString() string {
	var ss string

	switch t := any(s).(type) {
	case List[byte]:
		ss = string(t)
	}

	return ss
}
