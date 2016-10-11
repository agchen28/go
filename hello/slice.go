package main

import "fmt"

var months = [...]string{1: "一月", 2: "二月", 3: "三月", 12: "十二月"}

func testSlice() (int, int) {
	s := months[1:3]
	for _, value := range s {
		fmt.Println(value)
	}
	return len(s), cap(s)
}
