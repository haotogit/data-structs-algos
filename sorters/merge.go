package sorters

import "../util"

type Merger struct {}

func (m *Merger) merge(list []interface{}) {
    mid := len(list)/2

    leftIdx := 0
    rightIdx := mid
    tmpList := make([]interface{}, len(list))
    copy(tmpList, list)

    k := 0
    for ; leftIdx < mid && rightIdx < len(list); k++ {
        if util.ElComparer(tmpList[leftIdx], tmpList[rightIdx]) == - 1 {
            list[k] = tmpList[leftIdx]
            leftIdx++
        } else {
            list[k] = tmpList[rightIdx]
            rightIdx++
        }
    }

    for leftIdx < mid {
        list[k] = tmpList[leftIdx]
        k++
        leftIdx++
    }

    for rightIdx < len(list) {
        list[k] = tmpList[rightIdx]
        k++
        rightIdx++
    }
}

func (m *Merger) SortIt(list []interface{}) {
    mid := (len(list)/2)
    if mid > 0 {
        m.SortIt(list[:mid])
        m.SortIt(list[mid:])
        m.merge(list)
    }
}

func NewMerger() *Merger {
    return &Merger{}
}
