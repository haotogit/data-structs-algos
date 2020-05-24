package sorters

import "../util"

type Quick struct {}

func (q *Quick) randPivotPart(list []string, first, last int) int {
    randn := util.GetRandIntn(first, last)
    tmp := list[randn]
    list[randn] = list[first]
    list[first] = tmp

    return q.partition(list, first, last)
}

func (q *Quick) partition(list []string, first, last int) int {
    pivItem := list[first]
    left := first
    right := last+1

    for {
        // find leftItem > pivItem
        left++
        for left < right && list[left] < pivItem {
            left++
        }

        // find rightItem < pivItem
        right--
        for right >= left && list[right] > pivItem {
            right--
        }

        if left >= right {
            break
        }

        tmp := list[left]
        list[left] = list[right]
        list[right] = tmp
    }

    list[first] = list[right]
    list[right] = pivItem
    return right
}

func (q *Quick) quickSort(list []string, first, last int) {
    if first < last {
        piv := q.partition(list, first, last)
        q.quickSort(list, first, piv-1)
        q.quickSort(list, piv+1, last)
    }
}

func (q *Quick) SortIt(list []string) {
    q.quickSort(list, 0, len(list)-1)
}

func NewQuick() *Quick {
    return &Quick{}
}
