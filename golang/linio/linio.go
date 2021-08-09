package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 1; i <= 100; i++ {
		fmt.Println(LinioTest(uint(i)))
	}

}

/**
 * LinioTest
 * maxNumber -> positive number
 * prints all numbers between 0 and maxNumber
 * if %3 then print Linio
 * if %5 then print IT
 * if %3 and %5 then print Linianos
 */
func LinioTest(input uint) string {
	first15 := getFirst15()
	resultForFirst15 := getResultOfFirst15()

	res := input % 15
	if first15[res] {
		return strconv.FormatUint(uint64(input), 10)
	}
	return resultForFirst15[res]
}

/**
 * getResultOfFirst15
 * return the first 15 strings results that we want
 * for memoization reasons
 */
func getResultOfFirst15() [15]string {
	return [15]string{
		"Linianos",
		"",
		"",
		"Linio",
		"",
		"IT",
		"Linio",
		"",
		"",
		"Linio",
		"IT",
		"",
		"Linio",
		"",
		"",
	}
}

/**
 * getFirst15
 * return the first 15 boolean results
 * that we want
 */
func getFirst15() [15]bool {
	return [15]bool{
		false,
		true,
		true,
		false,
		true,
		false,
		false,
		true,
		true,
		false,
		false,
		true,
		false,
		true,
		true,
	}
}
