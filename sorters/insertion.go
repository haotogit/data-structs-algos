package sorters

import (
    "../util"
)

type Insertioner struct {}
    // TODO
    //sortBy string
    //sortDesc bool

func (in *Insertioner) SortIt(list []interface{}) {
    for x := 1; x < len(list); x++ {
        y := x
        tmp := list[y]
        // if prevNode is is greater than nextNode
        // set nextNode to previous consecutively
        // until no more previous
        for ; y > 0 && util.ElComparer(list[y-1], tmp) == 1; y-- {
            list[y] = list[y-1]
        }

        // then set original next node to before since it was lesser
        list[y] = tmp
    }
}

func NewInsertioner() *Insertioner {
    return &Insertioner {}
}
