package main

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	if y1-y2 >= 1 {
		return (y1 - y2) * 10000
	} else if y1 < y2 {
		return 0
	}
	if m1 == m2 {
		if d1 > d2 {
			return (d1 - d2) * 15
		} else {
			return 0
		}
	} else if m1 > m2 {
		return (m1 - m2) * 500
	}
	return 0
}
