package main

import "fmt"

func (e employee) sayHi() {
	fmt.Println(e.Name)
}

func testMethod() {
	e := employee{1, "zzz", "1", "pzzz"}
	e.sayHi()
	fmt.Println(e.pname)
}
