package main

import (
	"fmt"
)

var list []int
var sum int

// Solver returns a two element array with the values that sum to target number (otherwise returns [2]int{0, 0})
// O(n) to iterate through list once, pushing elements each iteration
func solver(list []int, sum int) [2]int {
	dict := make(map[int]int)

	for _, v := range list {
		diff := sum - v
		if _, ok := dict[diff]; ok {
			return [2]int{diff, v}
		} else {
			dict[v] = diff
		}
	}

	return [2]int{0, 0}
}

func main() {
	list = []int{2, 7, 11, 15}
	sum = 9

	fmt.Println(list)

	fmt.Print("Here is the answer: ")
	fmt.Println(solver(list, sum))
}
