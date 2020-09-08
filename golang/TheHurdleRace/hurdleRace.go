package main

func findMax(a []int32) int32 {
	min := a[0]
	max := a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return max
}

// Complete the hurdleRace function below.
func hurdleRace(k int32, height []int32) int32 {
	checkHeight := findMax(height)
	if k > checkHeight {
		return 0
	}
	return checkHeight - k
}
