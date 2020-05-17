package sorters

type BaseSorterer interface {
    // TODO need to take in list of interface{}
    SortIt(list []string)
}

type BaseSorter struct {
    BaseSorterer
}

func SortererCriador(sortAlgo int) *BaseSorter {
    var newSorterer *BaseSorter
    switch sortAlgo {
        default:
            newSorterer = &BaseSorter{BubbleMaker()}
    }

    return newSorterer
}

func main() {

}
