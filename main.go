package main

import (
	"fmt"
	"sync"
)

// the number of the consumers +1, because a real programmer start to count from zero :)
const N int = 9+1

func producer(chList [N]chan []int){
	// create a matrix of data, every row matches a consumer
	var dataSend [N][]int
	for i:=0;i<(N*N);i++ {
		// a tricky way to assign every row for each consumer
		dataSend[i%N] = append(dataSend[i%N], i)
	}
	// send data to a specific part of the array that matches the producer
	for i:=0;i<N;i++ {
		chList[i] <- dataSend[i]
	}
}

func consumer(chList [N]chan []int, wg *sync.WaitGroup,i int){
	defer wg.Done()
	fmt.Printf("consumer%d value: %d\n", i, <-chList[i])
}

func main(){
	wg := sync.WaitGroup{}
	// I prefer to use a array of channel because do that every process has own
	var chList [N]chan []int
	// you have to initialize channels inside the channel array, to avoid nil value
	for i := range chList {
		chList[i] = make(chan []int)
	}

	go producer(chList)

	// launch n consumers
	for i:=0;i<N;i++ {
		wg.Add(1)
		go consumer(chList, &wg, i)
	}

	// wait only consumers because the program only need to wait that every channel are consumed
	wg.Wait()
}
