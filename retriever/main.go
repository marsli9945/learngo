package main

import (
	"fmt"
	"github.com/marsli9945/learngo/retriever/mock"
	real2 "github.com/marsli9945/learngo/retriever/real"
	"time"
)

const url = "http://www.imooc.com"

/**
表示任何类型：interface{}
Type assertion
Type switch
*/
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func downLoad(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})

	return s.Get(url)
}

/**
Type switch
*/
func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}

func main() {
	var r Retriever
	mockRetriever := mock.Retriever{"this is a fake www.imooc.com"}
	inspect(&mockRetriever)

	r = &real2.Retriever{
		"Mozilla/5.0",
		time.Minute,
	}
	inspect(r)

	// Type assertion 类型断言
	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println("Contents:", mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a Session")
	fmt.Println(session(&mockRetriever))

	//fmt.Println(downLoad(r))
}
