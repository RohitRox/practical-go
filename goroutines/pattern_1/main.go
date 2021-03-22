package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const WORKER_POOL_SIZE = 3

func Caplock(id int, ip <-chan string, op chan<- string) {
	for item := range ip {
		op <- strings.ToUpper(item)
		r := rand.Intn(100)
		time.Sleep(time.Millisecond * time.Duration(r))
		fmt.Printf("Done by Routine: %d\n", id)
	}
}

func main() {
	arr := []string{"deakin", "lancaster", "kent", "ulster", "penzance", "devon", "wrexham", "earley", "hull", "rhyl", "corby", "dundee", "oban"}
	size := len(arr)

	fmt.Printf("INPUT: %v Size: %d\n", arr, size)

	inputCh := make(chan string, size)
	outputCh := make(chan string, size)

	for i := 0; i < WORKER_POOL_SIZE; i++ {
		go Caplock(i, inputCh, outputCh)
	}

	for _, item := range arr {
		inputCh <- item
	}

	close(inputCh)

	var result []string

	var wg sync.WaitGroup
	wg.Add(size)

	go func() {
		for item := range outputCh {
			result = append(result, item)
			wg.Done()
		}
	}()

	fmt.Printf("Let's do something before wait\n")

	time.Sleep(time.Millisecond * 100)

	fmt.Printf("End wait\n")

	wg.Wait()

	fmt.Printf("FINAL: %v Size: %d\n", result, len(result))
}
