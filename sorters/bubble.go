package sorters

import (
    "strings"
)

type Bubbler struct {}

func (b *Bubbler) SortIt(list []string) {
    for x := 1; x < len(list) - 1; x++ {
        swapped := false
        for y := len(list) - 1; y >= x; y-- {
            tmp := list[y-1]
            if strings.Compare(tmp, list[y]) == 1 {
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
