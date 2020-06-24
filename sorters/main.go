package sorters

type SortMachina interface {
	Sort(list []interface{})
	less(a, b interface{}) bool
}

func sortCriador(sortAlg int) SortMachina {
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
	case 5:
		return NewHeaper(false)
	}

	return NewQuicker()
}
