package main

import (
	"math/rand"
	"testing"
)

var ref [1000000]int
var arr [1000000]int

func TestMergeSorts(t *testing.T) {
	// Setup
	for i := 0; i < len(ref); i++ {
		ref[i] = rand.Int()
	}

	// Async test
	for i := 0; i < len(arr); i++ {
		arr[i] = ref[i]
	}
	t.Run("Merge sort async", func(t *testing.T) {
		dummyChannel := make(chan int, 1)
		mergeSort(arr[:], 0, len(arr)-1, dummyChannel)
		<-dummyChannel

		if !isSliceSorted(arr[:]) {
			t.Errorf("Bad!")
		}
	})

	// Sync test
	for i := 0; i < len(arr); i++ {
		arr[i] = ref[i]
	}
	t.Run("Merge sort sync", func(t *testing.T) {
		mergeSortSync(arr[:], 0, len(arr)-1)

		if !isSliceSorted(arr[:]) {
			t.Errorf("Bad!")
		}
	})
}

func BenchmarkMergeSorts(b *testing.B) {
	// Async benchmark
	b.Run("Merge sort async", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			for i := 0; i < len(arr); i++ {
				arr[i] = rand.Int()
			}
			b.StartTimer()

			dummyChannel := make(chan int, 1)
			mergeSort(arr[:], 0, len(arr)-1, dummyChannel)
			<-dummyChannel
		}
	})

	// Sync benchmark
	b.Run("Merge sort sync", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			for i := 0; i < len(arr); i++ {
				arr[i] = rand.Int()
			}
			b.StartTimer()

			mergeSortSync(arr[:], 0, len(arr)-1)
		}
	})
}

func isSliceSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}
