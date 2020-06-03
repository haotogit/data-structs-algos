package sorters

type SortMaquina interface {
	Sort(list []interface{})
	less(a, b interface{}) bool
}

var AlgMap = map[int]string{
	0: "Bubble",
	1: "Selection",
	2: "Insertion",
	3: "Merge",
	4: "Quick",
	5: "HeapSort",
}

func SortererCriador(sortAlg int) SortMaquina {
	switch sortAlg {
	case 0:
		return NewBubbler()
	case 1:
		return NewSelectioner()
	case 2:
		return NewInsertioner()
	case 3:
		return NewMerger()
	case 4:
		return NewQuicker()
	}

	return NewQuicker()
}
