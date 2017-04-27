package main

import (
	"flag"
	"fmt"
	// "time"
)

var (
	FirstUrl = flag.String("url", "http://rileystrong.com", "The first URL to visit")
	NWorkers = flag.Int("n", 10, "The max number of concurrent workers")
	MaxUrls = flag.Int("m", 10, "The max URLs to visit before stopping.")
)

func worker(id int, queue <-chan *UrlNode, nextUrl chan<- *UrlNode, completions chan<- *UrlNode) {
    for urlNode := range queue {
		urlNode.Process()
		// Bubble up in the background so not blocking & cause worker deadlock
		// Must pass in by argument vs closure since worker's re-used
		go func(deferredUrlNode *UrlNode) {
			defer func() {completions <- deferredUrlNode}()
			for _, urlStr := range deferredUrlNode.LinkedUrls {
				new_u, err := NewUrlNode(urlStr)
				// fmt.Println("after NewUrlNode list", deferredUrlNode.LinkedUrls)
				// fmt.Println("enqueueing", new_u)
				if err == nil {
					nextUrl <- new_u // blocks until worker picks it up
				}
			}
		}(urlNode)
    }
}

func main() {
	flag.Parse()

	queue := make(chan *UrlNode, 100) // TODO(riley): not sure optimal size
	nextUrl := make(chan *UrlNode) // must be blocking to avoid race conditions
	completions := make(chan *UrlNode)

	// Both urlsVisited and enqueued and  probably could be managed with atomic
	// ints, but nice to do everything just with channels.
	enqueued := 0 // Ensure we keep waiting when URLs are in the queue
	urlsVisited := make(map[string]bool)
	urlsVisiting := make(map[string]bool)

	// Seed with first URL
	urlNode, err := NewUrlNode(*FirstUrl)
	if err != nil {
		panic(err) // First URL failed - quit
	}

	// TODO(riley): DRY this out. Dup'd below
	queue <- urlNode
	enqueued++ 
	urlsVisiting[urlNode.UrlString] = true

	// Start all the workers
	for w := 1; w <= *NWorkers; w++ {
		go worker(w, queue, nextUrl, completions)
	}

	fmt.Println("[") // Manual JSON :-/
	// Keep going so long as there are URLs to process or we reach the max
	for {
		done := false
		select {
			case u := <- nextUrl:
				_, exists := urlsVisiting[u.UrlString]
				if !exists && len(urlsVisited) < *MaxUrls {
					queue <- u
					enqueued++
					urlsVisiting[u.UrlString] = true
				}
			case completed := <- completions:
				enqueued--
				urlsVisited[completed.UrlString] = true
				if len(urlsVisited) > 1 {
					fmt.Print(",\n") // Proper JSON formatting
				}
				completed.PrintResults()
			default:
				done = (enqueued == 0)
		}
		if done || len(urlsVisited) >= *MaxUrls  {
			break
		}
	}
	fmt.Print("\n]\n") // Manual JSON :-/

	// Leave some time to show concurrent actions continuing for debugging
	// fmt.Println("exit")
	// time.Sleep(100 * time.Millisecond)
}
