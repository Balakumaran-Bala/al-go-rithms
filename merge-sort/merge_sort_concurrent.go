package main

import (
	"math/rand"
)

func main() {
	var arr [100]int

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Int()
	}

	ch := make(chan int, 1)
	mergeSort(arr[:], 0, len(arr)-1, ch)
}

func mergeSort(arr []int, start int, end int, quit chan int) {
	if start == end {
		quit <- 1
		return
	}
	middle := (start + end) / 2
	ch := make(chan int, 2)
	go mergeSort(arr, start, middle, ch)
	go mergeSort(arr, middle+1, end, ch)
	<-ch
	<-ch
	merge(arr, start, end)
	quit <- 1
}

func merge(arr []int, start int, end int) {
	middle := (start + end) / 2
	left := make([]int, middle-start+1)
	copy(left, arr[start:middle+1])
	right := make([]int, end-middle)
	copy(right, arr[middle+1:end+1])
	for i, j := 0, 0; i < len(left) || j < len(right); {
		if (j == len(right)) || (i < len(left) && left[i] < right[j]) {
			arr[start+i+j] = left[i]
			i++
		} else {
			arr[start+i+j] = right[j]
			j++
		}
	}
}
