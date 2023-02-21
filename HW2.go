package main

import (
	"fmt"
	"sync"
)

var intaverage int32
var intmin int32
var intmax int32

func main() {
	fmt.Println("please enter 5 values to be compared: ")
	var wg sync.WaitGroup
	myslice := make([]int32, 0)
	var input int32
	for i := 0; i < 5; i++ {
		fmt.Scan(&input)
		myslice = append(myslice, input)
	}
	wg.Add(3)                // add 3 processes to the wait group
	go average(myslice, &wg) // pass in the slice along with a reference to the wait group
	go max(myslice, &wg)     // the go keyword run the function in a new thread
	go min(myslice, &wg)
	wg.Wait() // waits for all processes to run wg.Done()
	fmt.Println("average: " + fmt.Sprint(intaverage))
	fmt.Println("max: " + fmt.Sprint(intmax))
	fmt.Println("min: " + fmt.Sprint(intmin))
}

func average(slice []int32, wg *sync.WaitGroup) {
	defer wg.Done() // this will decrement the wait group when the function completes
	var out int32
	for i := 0; i < len(slice); i++ {
		out += slice[i]
	}
	out /= int32(len(slice))
	intaverage = out
}

func max(slice []int32, wg *sync.WaitGroup) {
	defer wg.Done()
	var out int32 = slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] > out {
			out = slice[i]
		}
	}
	intmax = out
}

func min(slice []int32, wg *sync.WaitGroup) {
	defer wg.Done()
	var out int32 = slice[0]
	for i := 1; i < len(slice); i++ {
		if slice[i] < out {
			out = slice[i]
		}
	}
	intmin = out
}

