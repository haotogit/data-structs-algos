package sorters

import (
	"../util"
)

type Selectioner struct{}

func (s *Selectioner) Sort(list []interface{}) {
	for cursor := 0; cursor < len(list); cursor++ {
		low := list[cursor]
		prevIdx := cursor
		for x := cursor + 1; x < len(list); x++ {
			if s.less(list[x], low) {
				low = list[x]
				prevIdx = x
			}
		}

		tmp := list[prevIdx]
		list[prevIdx] = list[cursor]
		list[cursor] = tmp
	}
}

func (s Selectioner) less(a, b interface{}) bool {
	return util.Greater(a, b) == -1
}

func NewSelectioner() SortMachina {
	return &Selectioner{}
}
