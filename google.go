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
	go func() { c <- Image(query) }() // fan-in pattern
	go func() { c <- Video(query) }()

	timeout := time.After(80 * time.Millisecond) // timeout for the whole process

	for range 3 {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Taking so long")
			return
		}
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
