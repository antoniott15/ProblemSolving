package main

import "fmt"

func rotations(a []int32, d int32) {

	var newArray = make([]int32, len(a))
	for i := 0; i < len(a); i++ {
		var newD int32 = int32(i) - d
		if newD < 0 {
			newArray[newD+int32(len(a))] = a[i]
		} else {
			newArray[newD] = a[i]
		}
	}

	for _, elements := range newArray {
		fmt.Print(elements, " ")
	}
}
