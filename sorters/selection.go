package sorters

import (
    "strings"
)

type Selectioner struct {}

func (s *Selectioner) SortIt(list []string) {
    for cursor := 0; cursor < len(list); cursor++ {
        low := list[cursor]
        prevIdx := cursor
        for x := cursor+1; x < len(list); x++ {
            if strings.Compare(list[x], low) == - 1 {
                low = list[x]
                prevIdx = x
            }
        }

        tmp := list[prevIdx]
        list[prevIdx] = list[cursor]
        list[cursor] = tmp
    }
}

func NewSelectioner() *Selectioner {
    return &Selectioner{}
}
