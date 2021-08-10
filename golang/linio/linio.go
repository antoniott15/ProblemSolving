package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {

	MAX := 100

	// Serial way
	for i := 1; i <= MAX; i++ {
		fmt.Println(LinioTest(uint(i)))
	}

	fmt.Println("Same test but using goroutines")

	// Concurrent with waitgroups
	LinioConcurrentTestWaitGroups(uint(MAX))

}

/**
 * LinioConcurrentTestWaitGroups
 * maxNumber -> positive number
 * prints all numbers between 1 and maxNumber
 * if %3 then print Linio
 * if %5 then print IT
 * if %3 and %5 then print Linianos
 */
func LinioConcurrentTestWaitGroups(concurrentCalls uint) {
	var wg sync.WaitGroup
	wg.Add(int(concurrentCalls))

	first15 := getFirst15()
	resultForFirst15 := getResultOfFirst15()

	for i := 1; i <= int(concurrentCalls); i++ {
		go func(i int) {
			defer wg.Done()
			res := i % 15
			if first15[res] {
				fmt.Println(strconv.FormatUint(uint64(i), 10))
			}
			fmt.Println(resultForFirst15[res])
		}(i)
	}

	wg.Wait()
}

/**
 * LinioTest
 * maxNumber -> positive number
 * prints all numbers between 1 and maxNumber
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
