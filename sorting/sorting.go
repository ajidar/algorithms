// Implementations of various sorting algorithms.

package sorting

import (
	"math/rand"
)

// InsertionSort is an n^2 sorting algorithm where the the input is sorted from left to right.
// An element is is inserted into the appropriate place on the left side of the list, then we move
// on to the next element.
func InsertionSort(input []int) {
	if len(input) < 2 {
		return
	}

	// Start from the second element in the list.
	for i := 1; i < len(input); i++ {

		// Keep moving the element backwards (by swapping it with the previous element) until it is in the correct spot.
		for j := i; j > 0; j-- {
			if input[j] < input[j-1] {
				input[j], input[j-1] = input[j-1], input[j]
			} else {
				break
			}
		}
	}
}

// SelectionSort is an n^2 sorting algorithm where the input is sorted by continually finding the smallest element,
// and moving it to the left side of the list.
func SelectionSort(input []int) {
	if len(input) < 2 {
		return
	}

	// Everything between 0 and i is already sorted.
	for i := 0; i < len(input); i++ {

		// Starting from the first element in the unsorted part of the list, find the smallest element remaining.
		j := i
		smallest := j
		j++
		for j < len(input) {
			if input[j] < input[smallest] {
				smallest = j
			}

			j++
		}

		// Stick the next smallest element at the end of the already sorted portion.
		input[i], input[smallest] = input[smallest], input[i]
	}
}

// Quicksort is an implementation of the common in-place Quicksort algorithm.
// It proxies for an internal function which handles the actual sorting.
func Quicksort(input []int) {

	quicksortInternal(input, 0, len(input)-1)
}

// quicksortInternal handles the actual machinations of the Quicksort algorithm.
// It takes an unsorted array, and the start and end indexes of the current sub-array.
// It partitions the elements of the sub-array around a random pivot,
// and then recursively process the sub-arrays to the left and right of the pivot.
func quicksortInternal(input []int, start int, end int) {

	if end <= start {
		return
	}

	// Partition the current segment.
	pivotIndex := partition(input, start, end)

	// Sort the left of the pivot.
	quicksortInternal(input, start, pivotIndex-1)

	// Sort the right of the pivot.
	quicksortInternal(input, pivotIndex+1, end)
}

// partition is a utility function used by Quicksort.
// Given an array, a start index, and an end index, it partitions the elements in the array around a randomly
// chosen pivot.
// Lesser elements will be placed to the left of the pivot, others to the right.
// The index of the pivot is returned.
func partition(input []int, start int, end int) int {

	// Randomly choose a pivot and temporarily stash it at the front of the segment.
	pivotIndex := rand.Intn(end-start) + start
	pivot := input[pivotIndex]
	input[start], input[pivotIndex] = input[pivotIndex], input[start]

	// Track pointers at either end.
	left, right := start, end

	// Move the lesser elements to the left of the pivot position,
	// and the greater elements to the right.
	for left < right {

		// Find the leftmost larger element.
		for input[left] <= pivot {
			left++

			// Don't exceed the bounds of the segment.
			if left == right {
				break
			}
		}

		// Find the rightmost smaller element.
		for input[right] > pivot {
			right--
		}

		// Swap the values if we didn't cross over.
		if left < right {
			input[left], input[right] = input[right], input[left]
		}
	}

	// Put the pivot back in place.
	input[start], input[right] = input[right], input[start]

	// Return the index of the pivot.
	return right
}

// QuicksortSimple is a simplified, and less efficient, version of Quicksort.
// It is NOT done in place, and does not use a separate partition function.
// Like the in place Quicksort, the sorting algorithm operates by picking a pivot element and then placing the
// lesser elements to the left, and the greater elements to the right.
// The left and right sides are then sorted recursively.
// Each call creates 2 new lists and appends them together.
func QuicksortSimple(original []int) []int {

	if len(original) < 2 {
		return original
	}

	pivot := original[0]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 1; i < len(original); i++ {
		if original[i] < pivot {
			left = append(left, original[i])
		} else {
			right = append(right, original[i])
		}
	}

	left = QuicksortSimple(left)
	right = QuicksortSimple(right)
	left = append(left, pivot)

	return append(left, right...)
}

// MergeSort is a sorting algorithm based on the idea of repeatedly merging sorted lists.
// The input is continually split into smaller and smaller lists, until they becomes lists of one.
// These lists are then merged back together.
func MergeSort(unsorted []int) []int {

	if unsorted == nil || len(unsorted) < 2 {
		return unsorted
	}

	mid := len(unsorted) / 2
	left := unsorted[0:mid]
	right := unsorted[mid:]

	return merge(MergeSort(left), MergeSort(right))
}

// merge is a utility function used  by MergeSort to merge two sorted arrays.
func merge(left []int, right []int) []int {

	if len(left) == 0 {
		return right
	}

	if len(right) == 0 {
		return left
	}

	merged := make([]int, 0)

	lptr, rptr := 0, 0

	for lptr < len(left) || rptr < len(right) {

		if lptr == len(left) {
			return append(merged, right[rptr:]...)
		}

		if rptr == len(right) {
			return append(merged, left[lptr:]...)
		}

		if right[rptr] < left[lptr] {
			merged = append(merged, right[rptr])
			rptr++
		} else {
			merged = append(merged, left[lptr])
			lptr++
		}
	}

	return merged
}
