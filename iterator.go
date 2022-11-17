package vulture

type iterator[T any] struct {
	val List[T]
}

func (i iterator[T]) OldScanRight(init T, fn func(accumulator T, val T) T) iterator[T] {
	var list List[T]
	var acc T = init

	for n := range i.val {
		acc = fn(acc, i.val[len(i.val)-1-n])
		list = append(List[T]{acc}, list...)
	}

	return iterator[T]{list}
}

func (i iterator[T]) ScanRight(init T, fn func(accumulator T, val T) T) iterator[T] {
	scanRight(&i.val, init, fn)
	return iterator[T]{i.val}
}

func (i iterator[T]) ScanLeft(init T, fn func(accumulator, val T) T) iterator[T] {
	var list List[T]
	var acc T = init

	for _, v := range i.val {
		acc = fn(acc, v)
		list = append(list, acc)
	}

	return iterator[T]{list}
}

func (i iterator[T]) FoldRight(init T, fn func(accumulator, val T) T) T {
	var acc T = init

	for n := range i.val {
		acc = fn(i.val[len(i.val)-1-n], acc)
	}

	return acc
}

func (i iterator[T]) FoldLeft(init T, fn func(accumulator, val T) T) T {
	var acc T = init

	for _, v := range i.val {
		acc = fn(acc, v)
	}

	return acc
}

func (i iterator[T]) Filter(fn func(T) bool) iterator[T] {
	list := List[T]{}

	for _, v := range i.val {
		if fn(v) {
			list = append(list, v)
		}
	}

	return iterator[T]{list}
}

func (i iterator[T]) Collect() List[T] {
	return i.val
}

type iteratorRef[T any] struct {
	val *List[T]
}

// func (i iteratorRef[T]) ScanRight(init T, fn func(accumulator T, val T) T) iterator[T] {

// }

func scanRight[T any](list *List[T], init T, fn func(accumulator, val T) T) {
	var acc T = init

	reverse(list)

	ll := *list

	for i, n := range ll {
		acc = fn(acc, n)
		ll[i] = acc
	}

	reverse(list)
}
