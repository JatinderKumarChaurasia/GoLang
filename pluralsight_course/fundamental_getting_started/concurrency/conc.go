package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	var waitgrp sync.WaitGroup
	fmt.Println("In Concurrency Channel")
	waitgrp.Add(2)
	fmt.Println(time.Now())
	go func() {
		defer waitgrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()
	fmt.Println(time.Now())
	go func() {
		defer waitgrp.Done()
		fmt.Println("Shivani Sharma")
	}()

	//myChannel := make(chan int,5)

	waitgrp.Wait()
}
