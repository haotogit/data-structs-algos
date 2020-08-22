package sorters

import (
	"github.com/haotogit/data-structs-algos/util"
)

// parent k/2
// left 2k
// right 2k+1
type HeapMachina interface {
	SortMachina
	sink(k int)
	swap(a, b interface{})
}

type Heap struct {
	isMax bool
}

func (h Heap) Sort(list []interface{}) {
	n := len(list)

	for k := (n - 1) / 2; k >= 0; k-- {
		h.heepify(list, k, n-1)
	}

	for x := n - 1; x >= 0; x-- {
		h.swap(list, 0, x)
		h.heepify(list, 0, x)
	}
}

func (h Heap) heepify(list []interface{}, k, lim int) {
	if k >= lim {
		return
	}

	dest := h.getHiLo(list, k, lim)

	if k != dest {
		h.swap(list, k, dest)
		h.heepify(list, dest, lim)
	}
}

// build minHeap for desc
// build maxHeap for asc
func (h Heap) getHiLo(list []interface{}, k, lim int) int {
	dest := k * 2
	if dest >= lim {
		return k
	}

	if !h.isMax {
		if dest+1 < lim && h.less(list[dest+1], list[dest]) {
			dest++
		}

		if h.less(list[k], list[dest]) {
			dest = k
		}
	} else {
		if dest+1 < lim && h.less(list[dest], list[dest+1]) {
			dest++
		}

		if h.less(list[dest], list[k]) {
			dest = k
		}
	}

	return dest
}

func (h Heap) less(a, b interface{}) bool {
	return util.Greater(a, b) == -1
}

func (h Heap) swap(list []interface{}, i, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}

func NewHeaper(isMax bool) *Heap {
	return &Heap{isMax}
}
