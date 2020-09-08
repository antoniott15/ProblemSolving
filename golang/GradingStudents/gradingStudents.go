package main

import "math"

func nextFiveMult(a int32) int32 {
	var newValue float64 = float64(a) / 5
	newValue = math.Ceil(newValue)
	newValue = newValue * 5
	return int32(newValue)
}

func gradingStudents(grades []int32) []int32 {
	var newGrades []int32
	for _, elements := range grades {
		if elements >= 38 && 5-(elements%5) < 3 {
			newGrades = append(newGrades, nextFiveMult(elements))
		} else {
			newGrades = append(newGrades, elements)
		}

	}
	return newGrades
}
