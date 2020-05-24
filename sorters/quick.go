package sorters

//import "fmt"

type Quick struct {}

func (q *Quick) partition(list []string, first, last int) int {
    pivIdx := first
    pivItem := list[pivIdx]
    left := first
    right := last+1

    for {
        // find leftItem > pivItem
        left++
        for left <= last && list[left] < pivItem {
            left++
            if left == last {
                break
            }
        }

        // find rightItem < pivItem
        right--
        for right >= first && list[right] > pivItem {
            right--
            if right == first {
                break
            }
        }

        if left >= right {
            break
        }

        tmp := list[left]
        list[left] = list[right]
        list[right] = tmp
    }

    list[pivIdx] = list[right]
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
