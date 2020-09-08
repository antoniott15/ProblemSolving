package main

import (
	"sort"
)

func sortElem(elements map[int32]int) []int {
	var result []int

	for key, _ := range elements {
		result = append(result, int(key))
	}

	sort.Ints(result)

	return result
}
func pickingNumbers(a []int32) int32 {
	numbersRepeated := make(map[int32]int)

	for _, values := range a {
		if _, ok := numbersRepeated[values]; ok {
			numbersRepeated[values] = numbersRepeated[values] + 1
		} else {
			numbersRepeated[values] = 1
		}
	}

	reapeatedSorted := sortElem(numbersRepeated)
	max := 0
	for index, value := range reapeatedSorted {
		number1, _ := numbersRepeated[int32(value)]
		number2, _ := numbersRepeated[int32(reapeatedSorted[index])+1]

		if number1+number2 > max {
			max = number1 + number2
		}

	}
	return int32(max)
}
