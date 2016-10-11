package main

import "fmt"

func sum(values ...int) {
	total := 0
	for value := range values {
		total += value
	}
	fmt.Println(total)
}
