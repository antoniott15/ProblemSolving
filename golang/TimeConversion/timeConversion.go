package main

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {
	hour := s[0:2]
	min := s[3:5]
	second := s[6:8]
	apm := s[8:10]
	fmt.Println(apm)
	var newHour int
	if apm == "PM" {
		newHour, _ = strconv.Atoi(hour)
		if newHour < 12 {
			newHour += 12
		} else if newHour == 12 {
			newHour = newHour
		} else {
			newHour += 12
		}
		return (strconv.Itoa(newHour) + ":" + min + ":" + second)
	} else {
		newHour, _ = strconv.Atoi(hour)
		if newHour == 12 {
			return ("00:" + min + ":" + second)
		}
		return (hour + ":" + min + ":" + second)
	}
	return ""

}
