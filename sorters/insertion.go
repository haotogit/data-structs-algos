package sorters

import (
    "strings"
)

type insertioner struct {
    sortBy string
    sortDesc bool
}

func (in *insertioner) SortIt(list []string) {
    for x := 1; x < len(list); x++ {
        y := x
        tmp := list[y]
        // if prevNode is is greater than nextNode
        // set nextNode to previous consecutively
        // until no more previous
        for ; y > 0 && strings.Compare(list[y-1], tmp) == 1; y-- {
            list[y] = list[y-1]
        }

        // then set original next node to before since it was lesser
        list[y] = tmp
    }
}

func NewInsertioner(sortMethod, sortBy string, sortDesc bool) *BaseSorter {
    return &BaseSorter{
        &insertioner {
            sortBy,
            sortDesc,
        },
    }
}
