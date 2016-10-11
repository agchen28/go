package main

import "fmt"

func testArray(index int) (int, int) {
	a := [...]int{1, 2, 3}
	for _, value := range a {
		fmt.Println(value)

	}
	fmt.Println(a[index])
	return len(a), a[index]
}
