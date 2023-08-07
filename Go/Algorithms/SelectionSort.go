package algorithms

func SelectionSort(array []int) {
	/*
		Bubble sort, is a simple sorting algorithm that repeatedly steps through the input list element by element, comparing the current element with the one after it, swapping their values if needed.
		These passes through the list are repeated until no swaps had to be performed during a pass, meaning that the list has become fully sorted.
		The algorithm, which is a comparison sort, is named for the way the larger elements "bubble" up to the top of the list.

		Time Complexity: - Worst Case:   O(n^2) (quadratic)
						 - Average Case: O(n^2) (quadratic)
			             - Best Case:    O(n)   (linear) {Only if it's optimized for the best case scenario}


		Space Complexity: - All Cases:   O(1)   (constant) {In place}
	*/

	for i := range array { // Go through each index
		minimum := i // Set the index of the minimum element

		for j := range array[i+1:] { // Compare with the succeeding values
			if array[minimum] > array[j] { // If the current minimum element is greater than the element at index j, set the minimum element index to j
				minimum = j
			}
		}
		array[minimum], array[i] = array[i], array[minimum] // Swap the values, now the array is sorted upto and including the i'th index
	}
}
