package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web1   = fakeSearch("Web1")
	Image1 = fakeSearch("Image1")
	Video1 = fakeSearch("Video1")

	Web2   = fakeSearch("Web2")
	Image2 = fakeSearch("Image2")
	Video2 = fakeSearch("Video2")

	Web3   = fakeSearch("Web3")
	Image3 = fakeSearch("Image3")
	Video3 = fakeSearch("Video3")
)

type Search func(string) Result
type Result string

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s results for %q query", kind, query))
	}
}

func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }

	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}

func Google(query string) (results []Result) {
	c := make(chan Result)

	go func() { c <- First(query, Web1, Web2, Web3) }()
	go func() { c <- First(query, Image1, Image2, Image3) }() // fan-in pattern
	go func() { c <- First(query, Video1, Video2, Video3) }()

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
