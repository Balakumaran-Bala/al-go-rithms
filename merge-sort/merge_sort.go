package main

func mergeSortSync(arr []int, start int, end int) {
	if start == end {
		return
	}
	middle := (start + end) / 2
	mergeSortSync(arr, start, middle)
	mergeSortSync(arr, middle+1, end)
	mergeSync(arr, start, end)
}

func mergeSync(arr []int, start int, end int) {
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
