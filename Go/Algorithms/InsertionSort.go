package algorithms

func InsertionSort(array []int) {
	/*
	   Insertion sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a time by comparisons.
	   It is much less efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort.

	   Time Complexity: - Worst Case:    O(n^2) (quadratic)
	                    - Average Case:  O(n^2) (quadratic)
	                    - Best Case:     O(n)   (quadratic)

	   Space Complexity: - All Cases:    O(1)   (constant) {In place}
	*/

	for i := range array {
		for j := i - 1; j >= 0 && array[j] > array[j+1]; j-- {
			var temp = array[j]
			array[j] = array[j+1]
			array[j+1] = temp
		}
	}
}
