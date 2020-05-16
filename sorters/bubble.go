package sorters

import (
    "strings"
)

func Bubbler(list []string) {
    for x := 0; x < len(list) - 1; x++ {
        swapped := false
        for y := 0; y < len(list) - x - 1; y++ {
            tmp := list[y]
            if strings.Compare(tmp, list[y+1]) == 1 {
                list[y] = list[y+1] 
                list[y+1] = tmp
                swapped = true
            }
        }

        if !swapped {
            break
        }
    }
}
