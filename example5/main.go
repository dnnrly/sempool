package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	// Imagine this is the result of a search on a REST API
	results := []int{10, 15, 8, 3, 17, 20, 1, 6, 10, 9, 13, 19}

	var wg sync.WaitGroup
	sem := make(chan struct{}, 3)

	for _, d := range results {
		wg.Add(1)

		go func(wait int) {
			defer func() {
				wg.Done()

				// Release the semaphore resource
				<-sem
			}()

			// Aquire a single semaophore resource
			sem <- struct{}{}

			// Imagine this is a long running operation, perhaps another
			// REST API call
			log.Printf("Start waiting for %d seconds\n", wait)
			time.Sleep(time.Second * time.Duration(wait))
			log.Printf("Finished waiting for %d seconds\n", wait)
		}(d)
	}

	wg.Wait()

	log.Printf("Total time taken: %s\n", time.Now().Sub(start))
}
