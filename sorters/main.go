package sorters

type BaseSortererr interface {
    // TODO need to take in list of interface{}
    SortIt(list []string)
}

type BaseSorter struct {
    // TODO how exactly is this working?
    BaseSortererr
}

var AlgMap = map[int]string{
    0: "Bubble",
    1: "Insertion",
    2: "Selection",
    3: "Merge",
    4: "Quick",
    5: "Heap",
}

func SortererCriador(sortAlg int, sortBy string, sortDesc bool) *BaseSorter {
    if sortAlg == 0 {
        return NewBubbler(AlgMap[sortAlg], sortBy, sortDesc)
    } else if sortAlg == 1 {
        //return &BaseSorter{ NewInserter(AlgMap[sortAlg], sortBy, sortDesc) }
    }

    return  nil
}
