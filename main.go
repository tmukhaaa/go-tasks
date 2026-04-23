package main

import (
	"fmt"
	"time"
)

// Проблема: есть писатель но нет читателя и запись в закрытый канал
// Решение: использовать буферезированный канал или поставить писателя после читателя
func main() {
	ch := make(chan bool)
	//ch <- true
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- true
	time.Sleep(1 * time.Second)
}
