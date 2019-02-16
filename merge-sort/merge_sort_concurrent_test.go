package main

import (
	"math/rand"
	"os"
	"testing"
)

var arr [1000000]int
var arrsync [1000000]int

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Int()
		arrsync[i] = arr[i]
	}
}

func TestMergeSort(t *testing.T) {
	dummyChannel := make(chan int, 1)

	mergeSort(arr[:], 0, len(arr)-1, dummyChannel)
	<-dummyChannel

	if !isSliceSorted(arr[:]) {
		t.Errorf("Bad!")
	}
}

func TestMergeSortSync(t *testing.T) {
	mergeSortSync(arrsync[:], 0, len(arr)-1)

	if !isSliceSorted(arrsync[:]) {
		t.Errorf("Bad!")
	}
}

func isSliceSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
