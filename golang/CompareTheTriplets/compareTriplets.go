package main

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {
	var pointsA int32 = 0
	var pointsB int32 = 0

	for i := range a {
		if a[i] < b[i] {
			pointsB += 1
		} else if a[i] > b[i] {
			pointsA += 1
		} else {
			continue
		}
	}
	newVector := []int32{pointsA, pointsB}

	return newVector
}
