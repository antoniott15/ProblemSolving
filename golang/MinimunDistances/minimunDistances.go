package main

import (
	"math"
	"sort"
)

func minimumDistances(a []int32) int32 {

	getPairs := make(map[int32][]int32)
	for index, elements := range a {
		getPairs[elements] = append(getPairs[elements], int32(index))
	}
	var countIfHave = 0
	var newArr []int
	for _, elements := range getPairs {
		if len(elements) == 2 {
			countIfHave += 1
			newArr = append(newArr, int(math.Abs(float64(elements[0]-elements[1]))))
		}
	}
	if countIfHave == 0 {
		return -1
	}
	sort.Ints(newArr)

	return int32(newArr[0])
}
