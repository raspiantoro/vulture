package vulture

type ScanType int

const (
	ScanTypeMin ScanType = iota
	ScanTypeMax
)

type numbers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type hofable interface {
	numbers | string
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

func getMin[T hofable](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func getMax[T hofable](a, b T) T {
	if a < b {
		return b
	}

	return a
}

func (s List[T]) ScanLeft(types ScanType) List[T] {
	var ss List[T]
	prev := s[0]

	ops := getMin[T]
	if types == ScanTypeMax {
		ops = getMax[T]
	}

	for _, v := range s {
		toAppend := ops(prev, v)
		ss = append(ss, toAppend)
		prev = toAppend
	}

	return ss
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
