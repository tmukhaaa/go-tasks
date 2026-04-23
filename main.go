package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func unpredictableFunc() int {
	time.Sleep(time.Second * 2)
	return 42
}

func predictableFunc() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	ch := make(chan int)
	go func() { ch <- unpredictableFunc() }()
	select {
	case v := <-ch:
		return v, nil
	case <-ctx.Done():
		return 0, errors.New("timed out")
	}
}

func main() {
	value, err := predictableFunc()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}
