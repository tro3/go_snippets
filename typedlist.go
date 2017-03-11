package snippets

import "sort"

type TYPEAList []TYPEA

func (x *TYPEAList) Push(val TYPEA) {
	*x = append(*x, val)
}

func (x *TYPEAList) Pop() TYPEA {
	y := (*x)[len(*x)-1]
	*x = (*x)[0 : len(*x)-1]
	return y
}

func (x *TYPEAList) Insert(ind int, val TYPEA) {
	tmp := (*x)[ind:]
	*x = (*x)[0:ind]
	*x = append(*x, val)
	*x = append(*x, tmp...)
}

func (x TYPEAList) Filter(fn func(TYPEA) bool) TYPEAList {
	y := TYPEAList{}
	for _, val := range x {
		if fn(val) {
			y = append(y, val)
		}
	}
	return y
}

// Sorting support, given a Less function

type TYPEASort struct {
	list TYPEAList
	less func(a, b TYPEA) bool
}

func (x TYPEASort) Len() int {
	return len(x.list)
}

func (x TYPEASort) Swap(i, j int) {
	x.list[i], x.list[j] = x.list[j], x.list[i]
}

func (x TYPEASort) Less(i, j int) bool {
	return x.less(x.list[i], x.list[j])
}

func (x *TYPEAList) Sort(less func(a, b TYPEA) bool) {
	ts := TYPEASort{*x, less}
	sort.Sort(ts)
}
