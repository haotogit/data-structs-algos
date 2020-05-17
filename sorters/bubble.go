package sorters

import (
    "strings"
)

type Bubbler struct{
}

func BubbleMaker() *Bubbler {
    return &Bubbler{}
}

func (b *Bubbler) SortIt(list []string) {
    for x := 0; x < len(list) - 1; x++ {
        swapped := false
        for y := 1; y < (len(list) - x); y++ {
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
