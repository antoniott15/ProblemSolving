package main

func camelcase(s string) int32 {
	lower := ([]byte)(s)
	var count int32 = 1
	for _, elements := range lower {
		if elements <= 90 {
			count++
		}
	}
	return count
}
