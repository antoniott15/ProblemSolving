package main

import "sort"

func migratoryBirds(arr []int32) int32 {
	var counter = make(map[int32]int32)
	var i int32 = 1
	for _, values := range arr {
		counter[values] = counter[values] + i
	}
	keys := make([]int, 0, len(counter))
	var max int32 = 0
	for k := range counter {
		if counter[int32(k)] > max {
			max = counter[int32(k)]
		}
		keys = append(keys, int(k))
	}

	sort.Slice(keys, func(i, j int) bool { return counter[int32(keys[i])] > counter[int32(keys[j])] })
	var preselectec []int32
	for _, k := range keys {
		if counter[int32(k)] == max {
			preselectec = append(preselectec, int32(k))
		}
	}

	sort.Slice(preselectec, func(i, j int) bool { return int(preselectec[i]) < int(preselectec[j]) })
	return preselectec[0]
}
