package sorters

import (
	"../util"
)

type Insertioner struct{}

// TODO
//sortBy string
//sortDesc bool

func (in *Insertioner) Sort(list []interface{}) {
	for x := 1; x < len(list); x++ {
		y := x
		tmp := list[y]
		// if prevNode is is greater than nextNode
		// set nextNode to previous consecutively
		// until no more previous
		for ; y > 0 && in.less(list[y-1], tmp); y-- {
			list[y] = list[y-1]
		}

		// then set original next node to before since it was lesser
		list[y] = tmp
	}
}

func (in *Insertioner) less(a, b interface{}) bool {
	return util.ElComparer(a, b) == -1
}

func NewInsertioner() SortMaquina {
	return &Insertioner{}
}
