package main

func saveThePrisoner(n int32, m int32, s int32) int32 {
	a := s + m - 1
	if a > n {
		if a%n == 0 {
			return n
		}
		return a % n
	}

	return a
}
