package main

func minimumSwaps(arr []int32) int32 {

	var count = 0
	for i := 0; i < len(arr); i++ {
		for arr[i] != int32(i+1) {
			temp := arr[i]
			arr[i] = arr[temp-1]
			arr[temp-1] = temp
			count++
		}
	}
	return int32(count)
}
