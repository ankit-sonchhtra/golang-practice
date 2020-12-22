package main

import ( 
	"fmt"
)

func main() {
	xi := Xi{4, 5, 1, 3, 2}
	sort(xi)
}

func sort(s Sorter) {
	for i := 0; i < s.Len(); i++ {
		for j := 0; j < s.Len()-1; j++ {
			if s.Less(j+1, j) {
				s.Swap(j+1, j)
			}
		}
	}
	fmt.Println(s)
}
