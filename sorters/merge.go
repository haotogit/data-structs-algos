package sorters

type Merger struct {}

func merge(list []string) {
    mid := len(list)/2

    leftIdx := 0
    rightIdx := mid
    tmpList := append([]string{}, list...)

    k := 0
    for ; leftIdx < mid && rightIdx < len(list); k++ {
        if tmpList[leftIdx] < tmpList[rightIdx] {
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

func mergeSort(list []string) {
    mid := (len(list)/2)
    if mid > 0 {
        mergeSort(list[:mid])
        mergeSort(list[mid:])
        merge(list)
    }
}

func (m *Merger) SortIt(list []string) {
    mergeSort(list)
}

func NewMerger() *Merger {
    return &Merger{}
}
