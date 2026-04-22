package main

import (
	"fmt"
	"sync"
)

// проблема: программа заканчивалась до начала работы горутины
// решение: использование sync.WaitGroup
func main() {
	var wg sync.WaitGroup
	a := 5000
	wg.Add(a)
	for i := 0; i < a; i++ {
		go func(v int) {
			defer wg.Done()
			fmt.Println(v)
		}(i)
	}
	wg.Wait()
}
