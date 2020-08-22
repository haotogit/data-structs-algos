package sorters

import "github.com/haotogit/data-structs-algos/util"

type QuickerMachina interface {
	SortMachina
	partition(list []interface{}, first, last int) int
	quickSort(list []interface{}, first, last int)
}

type Quicker struct{}

func (q *Quicker) partition(list []interface{}, first, last int) int {
	pivIdx := last
	pivItem := list[pivIdx]
	// last as pivot first-1, first as pivot last+1
	left := first - 1
	right := last

	for {
		// find leftItem > pivItem
		left++
		for left < right && q.less(pivItem, list[left]) {
			left++
		}

		// find rightItem < pivItem
		right--
		for right > left && !q.less(pivItem, list[right]) {
			right--
		}

		if left >= right {
			break
		}

		tmp := list[left]
		list[left] = list[right]
		list[right] = tmp
	}

	// last as pivot swap for lesser item now at left
	// first as pivot swap for greater item now at right
	list[pivIdx] = list[left]
	list[left] = pivItem
	return left
}

func (q *Quicker) quickSort(list []interface{}, first, last int) {
	if first < last {
		piv := q.partition(list, first, last)
		q.quickSort(list, first, piv-1)
		q.quickSort(list, piv+1, last)
	}
}

func (q *Quicker) Sort(list []interface{}) {
	q.quickSort(list, 0, len(list)-1)
}

func (q Quicker) less(a, b interface{}) bool {
	return util.Greater(a, b) == -1
}

func NewQuicker() *Quicker {
	return &Quicker{}
}
