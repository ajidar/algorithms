// Tests for sorting algorithms.
// All tests run against randomly generated 500000 element arrays with a max of 1000000.
// Algorithm timings (ms) are also outputed.

package testing

import (
	"testing"
	"math/rand"
	"time"
	"fmt"
	"github.com/ajidar/algorithms/sorting"
)

// generateRandomArray returns an unsorted array of random integers, of the requested length.
// Elements will be random integers between 0 and max.
func generateRandomArray(length int, max int) []int {
	rand.Seed(time.Now().UTC().UnixNano())
	unsorted := make([]int, length)
	for i := 0; i < length; i++ {
		unsorted[i] =  rand.Intn(max)
	}

	return unsorted
}

func TestMergeSort(t *testing.T) {
	unsorted := generateRandomArray(500000, 1000000)

	start := time.Now()
	sorted := sorting.MergeSort(unsorted)
	elapsed := time.Since(start).Nanoseconds()/1000000

	for i := 1; i < len(sorted); i++ {
		if sorted[i] < sorted[i - 1] {
			t.Fatalf("Output is not sorted. %d is less than %d.", sorted[i], sorted[i - 1])
		}
	}

	fmt.Printf("Merge Sort: %d ms\n", elapsed)
}

func TestQuicksort(t *testing.T) {
	unsorted := generateRandomArray(500000, 1000000)

	start := time.Now()
	sorted := sorting.Quicksort(unsorted)
	elapsed := time.Since(start).Nanoseconds()/1000000

	for i := 1; i < len(sorted); i++ {
		if sorted[i] < sorted[i - 1] {
			t.Fatalf("Output is not sorted. %d is less than %d.", sorted[i], sorted[i - 1])
		}
	}

	fmt.Printf("Quicksort: %d ms\n", elapsed)
}

func TestQuicksortInPlace(t *testing.T) {
	unsorted := generateRandomArray(500000, 1000000)

	start := time.Now()
	sorting.QuicksortInPlace(unsorted)
	elapsed := time.Since(start).Nanoseconds()/1000000

	for i := 1; i < len(unsorted); i++ {
		if unsorted[i] < unsorted[i - 1] {
			t.Fatalf("Output is not sorted. %d is less than %d.", unsorted[i], unsorted[i - 1])
		}
	}

	fmt.Printf("Quicksort In Place: %d ms\n", elapsed)
}
