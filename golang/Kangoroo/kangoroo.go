package main

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if x1 == x2 && v1 == v2 {
		return "YES"
	} else {
		var x, y int32 = 0, 0
		for i := 1; i < 10000; i++ {
			x = v1 * int32(i)
			y = v2 * int32(i)
			if x1+x == x2+y {
				return "YES"
			}
		}
		return "NO"
	}
	return "NO"
}
