package main

import "net/http"

//Reader 测试用
type Reader interface {
	Read(p []byte) (n int, err error)
}

//Closer 测试用
type Closer interface {
	Close() error
}

//ReadWriter 测试用
type ReadWriter interface {
	Reader
}

//ReadWriteCloser 测试用
type ReadWriteCloser interface {
	Reader
	Closer
}

//Handler 接口处理函数
type Handler interface {
	Do(w http.ResponseWriter, req *http.Request) string
}
