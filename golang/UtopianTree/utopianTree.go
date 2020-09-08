package main

func utopianTree(n int32) int32 {
	var utopia = 0
	for i := 0; i < int(n); i++ {
		if i == 0 {
			utopia = 1
		}
		if i%2 == 1 {
			utopia = utopia + 1
		} else {
			utopia = utopia * 2
		}
	}
	if n == 0 {
		utopia = 1
	}
	return int32(utopia)

}
