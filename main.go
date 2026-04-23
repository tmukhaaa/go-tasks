package main

import (
	"fmt"
	"sync"
	"time"
)

// проблема: обычная мапа не потоконебезопасна, поэтому выявляется гонка данных
// решение: использовать sync.Map или мьютексы
func main() {
	x := sync.Map{}
	go func() { x.Store(1, 2) }()
	go func() { x.Store(1, 7) }()
	go func() { x.Store(1, 10) }()
	time.Sleep(100 * time.Millisecond)
	v, _ := x.Load(1)
	fmt.Println("x[1] =", v)
}
