package main

func twoStrings(s1 string, s2 string) string {
	checkIf := make(map[string]bool)
	for _, elements := range s2 {
		checkIf[string(elements)] = true
	}
	for _, elements := range s1 {
		if _, ok := checkIf[string(elements)]; ok {
			return "YES"
		}
	}
	return "NO"
}
