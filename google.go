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
	results = make([]Result, 3)

	webResult := Web(query)
	results = append(results, webResult)

	results = append(results, Image(query))

	results = append(results, Video(query))

	return
}
func main() {
	start := time.Now()
	golangQuery := Google("golang")

	timeElapsed := time.Since(start)

	fmt.Println(golangQuery)
	fmt.Println(timeElapsed)
}
