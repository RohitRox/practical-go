package main

import (
	"fmt"
	"math/rand"
	"strings"
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

	for {
		if size == 0 {
			fmt.Println("Wait is over")
			break
		}
		result = append(result, <-outputCh)
		size--
	}

	fmt.Println("Just did some work")

	fmt.Printf("FINAL: %v Size: %d\n", result, len(result))
}
