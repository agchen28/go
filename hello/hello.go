package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/user/testPackage"
)

func main() {
	pool := poolInit()
	c := pool.Get()
	defer c.Close()
	reply, err := redis.String(c.Do("HGET", "PC:ChaoGaoFan:Brand_0", "489"))
	if err == nil {
		fmt.Print(reply)
	} else {
		fmt.Print(err)
	}
	mux := http.NewServeMux()
	s := new(myServer)
	mux.Handle("/get", http.HandlerFunc(s.get))
	mux.Handle("/post", http.HandlerFunc(s.post))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))

	// var i testInt
	var testI interface{}
	testI = 2
	// i = 1
	sum(1, 2, 3, 4)
	fmt.Println(testI)
	testMethod()
	// testTemp()
	// testStruct()
	// testJSON()
	// testMap()
	// l, c := testSlice()
	// fmt.Println(l)
	// fmt.Println(c)
	testPackage.TestPackage()
	// lenth, _ := testArray(1)
	// fmt.Println(lenth)
	// medals := []string{"silver", "gold"}
	// for _, value := range medals {
	// 	fmt.Println(value)
	// }
	// x := 1
	// p := &x
	// fmt.Println(*p)
	// *p = 2
	// fmt.Println(x)
	// x = 3
	// fmt.Println(*p)
	// fmt.Println(len(medals))
}

type testInt int

// import (
// 	"flag"
// 	"fmt"
// 	"strings"
// )

// var n = flag.Bool("n", false, "omit trailing newline")
// var sep = flag.String("s", " ", "separator")

// func main() {
// 	flag.Parse()
// 	fmt.Print(strings.Join(flag.Args(), *sep))
// 	if !*n {
// 		fmt.Println()
// 	}
// }

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	if len(files) == 0 {
// 		countLines(os.Stdin, counts)
// 	} else {
// 		for _, arg := range files {
// 			f, err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 				continue
// 			}
// 			countLines(f, counts)
// 			f.Close()
// 		}
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }
// func countLines(f *os.File, counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// 	// NOTE: ignoring potential errors from input.Err()
// }

// import (
// 	"fmt"

// 	"github.com/user/stringutil"
// )

// func main() {

// 	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
// 	i, j := 123, 1
// 	k := i + j
// 	s := string(k)
// 	fmt.Printf(s)
// 	if i := 0; i > 0 {
// 		fmt.Printf("大于0")
// 	}
// }
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		if input.Text() == "" {
// 			break
// 		}
// 		counts[input.Text()]++
// 	}
// 	// NOTE: ignoring potential errors from input.Err()
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }
