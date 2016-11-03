package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
)

type myServer struct {
	name string
}

func (s myServer) get(w http.ResponseWriter, req *http.Request) {
	result := make(chan string)
	go func() {
		db, err := sql.Open("mssql", "server=192.168.1.221;user id=sa;password=123456;database=fanhuansqlserver")
		if err == nil {
			rows, err := db.Query("SELECT top 1 Id,Name FROM ChannelCode")
			if err == nil {
				defer rows.Close()
				for rows.Next() {
					var id int
					var name string
					err = rows.Scan(&id, &name)
					result <- name
				}
			} else {
				fmt.Println("Cannot connect: ", err.Error())
			}
		} else {
			fmt.Println("Cannot connect: ", err.Error())
		}
	}()

	fmt.Fprintf(w, <-result)
}

func (s myServer) post(w http.ResponseWriter, req *http.Request) {
	resultChan := make(chan string)
	// go func() {
	// 	resultChan <- get("489")
	// }()
	// fmt.Fprintf(w, <-resultChan)
	go func() {
		resultChan <- "post"
	}()
	fmt.Fprintf(w, <-resultChan)
}
