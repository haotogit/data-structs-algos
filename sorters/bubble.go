package sorters

import (
    "../util"
)

type Bubbler struct {}

func (b *Bubbler) SortIt(list []interface{}) {
    for x := 1; x < len(list) - 1; x++ {
        swapped := false
        for y := len(list) - 1; y >= x; y-- {
            tmp := list[y-1]
            if util.ElComparer(tmp, list[y]) == 1 {
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

func NewBubbler() *Bubbler {
    return &Bubbler{}
}
