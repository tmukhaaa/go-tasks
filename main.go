package main

import (
	"fmt"
	"sync"
)

// Проблема: горутины одновременно пытаются записать в счетчик, из за чего создается гонка данных
// и некоторые итерации прожевываются
// Решение: использование мьютаекса или атомика
func main() {
	var counter int64
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			//atomic.AddInt64(&counter, 1)
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
