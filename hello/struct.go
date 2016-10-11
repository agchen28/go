package main

import "fmt"

type employee struct {
	ID    int
	Name  string `json:"name"`
	Tag   string `json:"Tag,omitempty"`
	pname string
}

func testStruct() {
	var e employee
	e.ID = 1
	e.Name = "zzz"
	fmt.Println(e.Name)
}
