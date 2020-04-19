package main

import (
	"fmt"
	"math/rand"
	"time"
)

func FindRandomNumber(value int) {

	count := 1
	numberFound := false

	for {
		num := rand.Intn(1000)
		if num == value {
			numberFound = true
			break
		}
		count++
	}
	if numberFound {
		fmt.Printf("Number #%v found after %v Steps", value, count)
	}
}

func main() {
	// InfiniteLooping.go
	rand.Seed(time.Now().UnixNano())
	go FindRandomNumber(rand.Intn(100)) // using goroutine
	time.Sleep(5 * time.Second)
}
