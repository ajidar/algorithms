package sorting

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// generateRandomArray returns an unsorted array of random integers, of the requested length.
// Elements will be random integers between 0 and max.
func generateRandomArray(length int, max int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	unsorted := make([]int, length)
	for i := 0; i < length; i++ {
		unsorted[i] = rand.Intn(max)
	}

	return unsorted
}

func TestMergeSort(t *testing.T) {
	unsorted := generateRandomArray(500000, 1000000)

	start := time.Now()
	sorted := MergeSort(unsorted)
	elapsed := time.Since(start).Nanoseconds() / 1000000

	for i := 1; i < len(sorted); i++ {
		if sorted[i] < sorted[i-1] {
			t.Fatalf("Output is not sorted. %d is less than %d.", sorted[i], sorted[i-1])
		}
	}

	fmt.Printf("Merge Sort: %d ms\n", elapsed)
}

func TestQuicksort(t *testing.T) {
	unsorted := generateRandomArray(500000, 1000000)

	start := time.Now()
	sorted := Quicksort(unsorted)
	elapsed := time.Since(start).Nanoseconds() / 1000000

	for i := 1; i < len(sorted); i++ {
		if sorted[i] < sorted[i-1] {
			t.Fatalf("Output is not sorted. %d is less than %d.", sorted[i], sorted[i-1])
		}
	}

	fmt.Printf("Quicksort: %d ms\n", elapsed)
}

func TestQuicksortInPlace(t *testing.T) {
	unsorted := generateRandomArray(500000, 1000000)

	start := time.Now()
	QuicksortInPlace(unsorted)
	elapsed := time.Since(start).Nanoseconds() / 1000000

	for i := 1; i < len(unsorted); i++ {
		if unsorted[i] < unsorted[i-1] {
			t.Fatalf("Output is not sorted. %d is less than %d.", unsorted[i], unsorted[i-1])
		}
	}

	fmt.Printf("Quicksort In Place: %d ms\n", elapsed)
}

func TestInsertionSort(t *testing.T) {

	fmt.Println("Warning: Starting insertion sort. This could take a while.")

	unsorted := generateRandomArray(500000, 1000000)

	start := time.Now()
	InsertionSort(unsorted)
	elapsed := time.Since(start).Nanoseconds() / 1000000

	for i := 1; i < len(unsorted); i++ {
		if unsorted[i] < unsorted[i-1] {
			t.Fatalf("Output is not sorted. %d is less than %d.", unsorted[i], unsorted[i-1])
		}
	}

	fmt.Printf("Insertion Sort: %d ms\n", elapsed)
}

func TestSelectionSort(t *testing.T) {

	fmt.Println("Warning: Starting selection sort. This could take a while.")

	unsorted := generateRandomArray(500000, 1000000)

	start := time.Now()
	SelectionSort(unsorted)
	elapsed := time.Since(start).Nanoseconds() / 1000000

	for i := 1; i < len(unsorted); i++ {
		if unsorted[i] < unsorted[i-1] {
			t.Fatalf("Output is not sorted. %d is less than %d.", unsorted[i], unsorted[i-1])
		}
	}

	fmt.Printf("Selection Sort: %d ms\n", elapsed)
}
