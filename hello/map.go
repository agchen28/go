package main

import "fmt"

var m = map[string]string{
	"a": "aa",
	"b": "bb",
}

func testMap() {
	for _, value := range m {
		fmt.Println(value)
	}
}
