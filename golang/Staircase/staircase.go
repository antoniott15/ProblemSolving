package main

import "fmt"

func staircase(n int) {

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j == n-1 {
				fmt.Print("#")
			} else if i+j >= n {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
