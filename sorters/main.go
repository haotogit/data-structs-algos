package sorters

type SortMaquina interface {
    // TODO need to take in list of interface{}
    SortIt(list []string)
}

var AlgMap = map[int]string{
    0: "Bubble",
    1: "Insertion",
    2: "Selection",
    3: "Merge",
    4: "Quick",
    5: "HeapSort",
}

func SortererCriador(sortAlg int) SortMaquina {
    switch sortAlg {
        case 0:
            return NewBubbler()
        case 1:
            return NewInsertioner()
        case 2:
            return NewSelectioner()
        case 3:
            return NewMerger()
        case 4:
            return NewQuicker()
    }

    return  nil
}
