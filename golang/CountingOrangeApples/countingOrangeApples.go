package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var values []int32
	for i := s; i <= t; i++ {
		values = append(values, i)
	}

	var countingApples int32 = 0
	for _, elements := range apples {
		if elements > 0 {
			if a+elements >= s && a+elements <= t {
				countingApples++
			}
		}
	}

	var countingOrange int32 = 0
	for _, elements := range oranges {
		if b+elements >= s && b+elements <= t {
			countingOrange++
		}
	}

	fmt.Println(countingApples)
	fmt.Println(countingOrange)

}
