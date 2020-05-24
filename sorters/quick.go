package sorters

type Quicker struct {}

func (q *Quicker) partition(list []string, first, last int) int {
    pivIdx := last
    pivItem := list[pivIdx]
    // last as pivot first-1, first as pivot last+1
    left := first-1
    right := last

    for {
        // find leftItem > pivItem
        left++
        for left < right && list[left] < pivItem {
            left++
        }

        // find rightItem < pivItem
        right--
        for right > left && list[right] > pivItem {
            right--
        }

        if left >= right {
            break
        }

        tmp := list[left]
        list[left] = list[right]
        list[right] = tmp
    }

    // last as pivot swap for lesser item now at left
    // first as pivot swap for greater item now at right
    list[pivIdx] = list[left]
    list[left] = pivItem
    return left
}

func (q *Quicker) quickSort(list []string, first, last int) {
    if first < last {
        piv := q.partition(list, first, last)
        q.quickSort(list, first, piv-1)
        q.quickSort(list, piv+1, last)
    }
}

func (q *Quicker) SortIt(list []string) {
    q.quickSort(list, 0, len(list)-1)
}

func NewQuicker() *Quicker {
    return &Quicker{}
}
