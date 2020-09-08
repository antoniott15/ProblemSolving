package main

import "math"

func digit(num int32, place int) int32 {
	r := int(num) % int(math.Pow(10, float64(place)))
	if0 := int(math.Pow(10, float64(place-1)))
	if if0 == 0 {
		return -1
	}
	return int32(r / if0)
}
