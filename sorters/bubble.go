package sorters

import (
	"../util"
)

type Bubbler struct{}

func (b *Bubbler) Sort(list []interface{}) {
	for x := 1; x < len(list); x++ {
		swapped := false
		for y := len(list) - 1; y >= x; y-- {
			tmp := list[y-1]
			if b.less(tmp, list[y]) {
				list[y-1] = list[y]
				list[y] = tmp
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func (b Bubbler) less(a, bb interface{}) bool {
	return util.Greater(a, bb) == -1
}

func NewBubbler() SortMaquina {
	return &Bubbler{}
}
