package main

import "sort"

func equalizeArray(arr []int32) int32 {

	newarr := make(map[int32]int)
	for _, elements := range arr {
		newarr[elements] = newarr[elements] + 1
	}

	var result []int
	for _, values := range newarr {
		result = append(result, values)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	result[0] = result[len(result)-1]
	result[len(result)-1] = 0
	result = result[:len(result)-1]

	var counter int32 = 0
	for _, elements := range result {
		counter += int32(elements)
	}
	return counter
}
