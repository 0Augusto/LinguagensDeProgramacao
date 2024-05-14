package main

import (
	"fmt"
	"runtime"
	"sync"
)

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]

	var wg sync.WaitGroup
	wg.Add(2)

	runtime.Gosched()

	var leftResult, rightResult []int

	go func() {
		leftResult = mergeSort(left)
		wg.Done()
	}()

	go func() {
		rightResult = mergeSort(right)
		wg.Done()
	}()

	wg.Wait()

	return merge(leftResult, rightResult)
}

func main() {
	arr := []int{5, 9, 3, 7, 2, 8, 6, 1, 4}
	fmt.Println("Unsorted array:", arr)
	arr = mergeSort(arr)
	fmt.Println("Sorted array:", arr)
}
