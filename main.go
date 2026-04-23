package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(3)

	go fetchFromDb(ctx, ch, &wg)
	go fetchFromAPI(ctx, ch, &wg)
	go fetchFromCache(ctx, ch, &wg)

	select {
	case v := <-ch:
		log.Println(v)
		cancel()
	case <-time.After(20 * time.Millisecond):
		log.Println("timeout")
		cancel()
	}
	wg.Wait()
}

func fetchFromCache(ctx context.Context, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.After(10 * time.Millisecond):
		result <- "fetch from cache"
	case <-ctx.Done():
		log.Println("fetch from cache CANCELED")
	}
}

func fetchFromDb(ctx context.Context, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.After(50 * time.Millisecond):
		result <- "fetch from DB"
	case <-ctx.Done():
		log.Println("fetch from DB CANCELED")
	}
}

func fetchFromAPI(ctx context.Context, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.After(100 * time.Millisecond):
		result <- "fetch from API"
	case <-ctx.Done():
		log.Println("fetch from API CANCELED")
	}
}
