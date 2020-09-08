package main

func alternatingCharacters(s string) int32 {
	counter := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			counter++
		}
	}
	return int32(counter)
}
