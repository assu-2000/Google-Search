package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("Web")
	Image = fakeSearch("Image")
	Video = fakeSearch("Video")
)

type Search func(string) Result
type Result string

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s results for %q query", kind, query))
	}
}

func Google(query string) (results []Result) {
	c := make(chan Result)

	go func() { r := Web(query); c <- r }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	for range 3 {
		result := <-c
		results = append(results, result)
	}

	return
}
func main() {
	start := time.Now()
	golangQuery := Google("golang")

	timeElapsed := time.Since(start)

	fmt.Println(golangQuery)
	fmt.Println(timeElapsed)
}
