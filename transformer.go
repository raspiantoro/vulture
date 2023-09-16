package vulture

type transIter[D any, S any] struct {
	iter iterator[S]
}

// since Iter already called from transformer
// when going back to transformer, it cannot refer to original List reference
// we should give new reference to prevent panic
func (t transIter[D, S]) Transform() transformer[D, S] {
	l := t.iter.Collect()
	return transformer[D, S]{
		list:    t.iter.Collect(),
		refList: &l,
	}
}

func (t transIter[D, S]) Add(addition S, fn func(addition, val S) S) transIter[D, S] {
	return transIter[D, S]{t.iter.Add(addition, fn)}
}

func (t transIter[_, S]) Collect() List[S] {
	return t.iter.Collect()
}

type transIterRef[D any, S any] struct {
	iter iteratorRef[S]
}

func (t transIterRef[D, S]) Transform() transformer[D, S] {
	return transformer[D, S]{
		list:    *t.iter.Collect(),
		refList: t.iter.Collect(),
	}
}

func (t transIterRef[D, S]) Add(addition S, fn func(addition, val S) S) transIterRef[D, S] {
	return transIterRef[D, S]{t.iter.Add(addition, fn)}
}

func (t transIterRef[_, S]) Collect() *List[S] {
	return t.iter.Collect()
}

type transformer[D any, S any] struct {
	list    List[S]
	refList *List[S]
}

func Transformer[D any, S any](list List[S]) transformer[D, S] {
	return transformer[D, S]{
		list: list,
	}
}

func (t transformer[D, S]) Iter() transIter[D, S] {
	return transIter[D, S]{
		iter: t.list.Iter(),
	}
}

func (t transformer[D, S]) IterRef() transIterRef[D, S] {
	return transIterRef[D, S]{
		iter: t.refList.IterRef(),
	}
}

func (t transformer[D, S]) Map(fn func(val S) D) List[D] {
	listD := List[D]{}
	listS := t.Iter().iter.cloneVal()

	for _, s := range listS {
		listD = append(listD, fn(s))
	}

	return listD
}
