package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func testJSON() {
	var es = []employee{
		{1, "zzz", "zzz", ""},
		{2, "XIXI", "", ""},
	}
	data, err := json.MarshalIndent(es, "", " ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var es1 []employee
	if err := json.Unmarshal(data, &es1); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(es1[0].Name)
}
