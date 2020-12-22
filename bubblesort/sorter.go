package main

type Sorter interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

type Xi []int

func (xi Xi) Len() int {
	return len(xi)
}

func (xi Xi) Swap(i, j int) {
	xi[i], xi[j] = xi[j], xi[i]
}

func (xi Xi) Less(i, j int) bool {
	return xi[i] < xi[j]
}
