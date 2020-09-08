package main

import "sort"

func birthdayCakeCandles(ar []int) int {
	sort.Ints(ar)
	maxValue := ar[len(ar)-1]
	var count int = 0
	for _, element := range ar {
		if element == maxValue {
			count++
		}
	}
	return count
}
