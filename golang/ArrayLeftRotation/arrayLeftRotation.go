package main

func rotLeft(a []int32, d int32) []int32 {
	result := make([]int32, len(a))
	for i := 0; i < len(a); i++ {
		if i-int(d) < 0 {
			rest := i - int(d)
			result[len(a)+rest] = a[i]
		} else {
			result[i-int(d)] = a[i]
		}
	}
	return result
}
