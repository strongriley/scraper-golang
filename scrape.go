package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	// "time"
)

var (
	FirstUrl = flag.String("url", "http://rileystrong.com", "The first URL to visit")
	NWorkers = flag.Int("n", 10, "The max number of concurrent workers")
	MaxUrls  = flag.Int("m", 10, "The max URLs to visit before stopping.")
)

func worker(id int, queue <-chan *UrlNode, completions chan<- *UrlNode) {
	for urlNode := range queue {
		urlNode.Process()
		// Bubble up in the background so not blocking & cause worker deadlock
		// Must pass in by argument vs closure since worker's re-used
		go func(deferredUrlNode *UrlNode) {
			completions <- deferredUrlNode
		}(urlNode)
	}
}

func main() {
	flag.Parse()

	queue := make(chan *UrlNode, 100) // TODO(riley): not sure optimal size
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
		go worker(w, queue, completions)
	}

	// Allow Ctrl-C to terminate the app with valid JSON still
	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt)

	fmt.Println("[") // Manual JSON :-/
	// Keep going so long as there are URLs to process or we reach the max
	for {
		done := false
		select {
			case <-interrupt:
				done = true
			default:
		}

		if done {
			break
		}

		select {
		case completed := <-completions:
			for _, urlStr := range completed.LinkedUrls {
				new_u, err := NewUrlNode(urlStr)
				if err != nil {
					continue
				}
				_, exists := urlsVisiting[new_u.UrlString]
				if !exists && len(urlsVisited) < *MaxUrls {
					queue <- new_u
					enqueued++
					urlsVisiting[new_u.UrlString] = true
				}
			}
			enqueued--
			urlsVisited[completed.UrlString] = true
			if len(urlsVisited) > 1 {
				fmt.Print(",\n") // Proper JSON formatting
			}
			completed.PrintResults()
		default:
			done = (enqueued == 0)
		}

		if done || len(urlsVisited) >= *MaxUrls {
			break
		}
	}
	fmt.Print("\n]\n") // Manual JSON :-/
}
