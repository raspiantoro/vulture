package vulture

type iterator[T any] struct {
	val List[T]
}

func (i iterator[T]) cloneVal() List[T] {
	list := make(List[T], len(i.val))
	copy(list, i.val)

	return list
}

func (i iterator[T]) ScanRight(init T, fn func(accumulator T, val T) T) iterator[T] {
	list := i.cloneVal()
	scanRight(&list, init, fn)

	return iterator[T]{list}
}

func (i iterator[T]) ScanLeft(init T, fn func(accumulator, val T) T) iterator[T] {
	list := i.cloneVal()
	scanLeft(&list, init, fn)

	return iterator[T]{list}
}

func (i iterator[T]) Add(addition T, fn func(addition, val T) T) iterator[T] {
	list := i.cloneVal()
	add(&list, addition, fn)

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

func (i iteratorRef[T]) Collect() *List[T] {
	return i.val
}

func (i iteratorRef[T]) ScanRight(init T, fn func(accumulator T, val T) T) iteratorRef[T] {
	scanRight(i.val, init, fn)
	return iteratorRef[T]{i.val}
}

func (i iteratorRef[T]) ScanLeft(init T, fn func(accumulator T, val T) T) iteratorRef[T] {
	scanLeft(i.val, init, fn)
	return iteratorRef[T]{i.val}
}

func (i iteratorRef[T]) Add(addition T, fn func(addition, val T) T) iteratorRef[T] {
	add(i.val, addition, fn)

	return iteratorRef[T]{i.val}
}

func (i iteratorRef[T]) FoldRight(init T, fn func(accumulator, val T) T) T {
	var acc T = init

	// for n := range *i.val {
	// 	acc = fn(i.val[len(i.val)-1-n], acc)
	// }

	return acc
}

func (i iteratorRef[T]) FoldLeft(init T, fn func(accumulator, val T) T) T {
	var acc T = init

	// for _, v := range *i.val {
	// 	acc = fn(acc, v)
	// }

	return acc
}

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

func scanLeft[T any](list *List[T], init T, fn func(accumulator, val T) T) {
	var acc T = init
	ll := *list

	for i, n := range ll {
		acc = fn(acc, n)
		ll[i] = acc
	}
}

func add[T any](list *List[T], addition T, fn func(addition, val T) T) {
	ll := *list

	for i, n := range ll {
		n = fn(addition, n)
		ll[i] = n
	}
}
