package main

import (
	"retriever/mock"
	"fmt"
	"retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever)  string {
	return r.Get("http://www.imooc.com")
}

func main() {
	var r  Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	inspect(r)
    r = &real.Retriever{
    	UserAgent:"Mozilla/5.0",
    	TimeOut:time.Minute,
	}

	inspect(r)
	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n",r,r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Conrents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UseAgent", v.UserAgent)
	}
}
