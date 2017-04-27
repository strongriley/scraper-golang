package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	FirstUrl = flag.String("url", "http://rileystrong.com", "The first URL to visit")
	NWorkers = flag.Int("n", 2, "The max number of concurrent workers")
	MaxUrls = flag.Int("m", 5, "The max URLs to visit before stopping. Not a strong guarantee")
)

func worker(id int, queue <-chan *UrlNode, nextUrl chan<- *UrlNode, completions chan<- *UrlNode) {
    for u := range queue {
		u.Process()
		// Bubble up in the background so not blocking & cause worker deadlock
		go func() {
			defer func() {completions <- u}()
			for _, urlStr := range u.linkedUrls {
				new_u, err := NewUrlNode(urlStr)
				if err == nil {
					nextUrl <- new_u // blocks until worker picks it up
				}
			}
		}()
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

	// Seed with first URL
	urlNode, err := NewUrlNode(*FirstUrl)
	if err != nil {
		panic(err) // First URL failed - quit
	}
	queue <- urlNode
	enqueued++ // TODO(riley): DRY this out. Dup'd below

	// Start all the workers
	for w := 1; w <= *NWorkers; w++ {
		go worker(w, queue, nextUrl, completions)
	}

	// Keep going so long as there are URLs to process or we reach the max
	for {
		done := false
		select {
			case u := <- nextUrl:
				// NB(riley): Race conditions may result in the same URL being
				// procesed twice
				// TODO(riley): fix race conditions
				_, exists := urlsVisited[u.url.String()]
				if !exists {
					queue <- u
					enqueued++
				} else {
					fmt.Println("skipping", u)
				}
			case u := <- completions:
				enqueued--
				urlsVisited[u.url.String()] = true
				fmt.Println(u.url.String())
				u.PrintResults()
			default:
				if enqueued == 0 {
					fmt.Println("Nothing left to do, ending")
					done = true
				}
		}
		if done || len(urlsVisited) >= *MaxUrls  {
			break
		}
	}

	// Leave some time to show concurrent actions continuing
	fmt.Println("exit")
	time.Sleep(100 * time.Millisecond)
}
