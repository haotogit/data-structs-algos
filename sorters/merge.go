package sorters

import "../util"

type Merger struct{}

func (m *Merger) merge(list []interface{}) {
	mid := len(list) / 2

	leftIdx := 0
	rightIdx := mid
	tmpList := make([]interface{}, len(list))
	copy(tmpList, list)

	k := 0
	for ; leftIdx < mid && rightIdx < len(list); k++ {
		if m.less(tmpList[rightIdx], tmpList[leftIdx]) {
			list[k] = tmpList[leftIdx]
			leftIdx++
		} else {
			list[k] = tmpList[rightIdx]
			rightIdx++
		}
	}

	for leftIdx < mid {
		list[k] = tmpList[leftIdx]
		k++
		leftIdx++
	}

	for rightIdx < len(list) {
		list[k] = tmpList[rightIdx]
		k++
		rightIdx++
	}
}

func (m *Merger) Sort(list []interface{}) {
	mid := (len(list) / 2)
	if mid > 0 {
		m.Sort(list[:mid])
		m.Sort(list[mid:])
		m.merge(list)
	}
}

func (m Merger) less(a, b interface{}) bool {
	return util.Greater(a, b) == -1
}

func NewMerger() SortMaquina {
	return &Merger{}
}
