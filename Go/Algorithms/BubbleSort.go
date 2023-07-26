package algorithms

func BubbleSort(array []int) {
	/*
		Bubble sort, is a simple sorting algorithm that repeatedly steps through the input list element by element, comparing the current element with the one after it, swapping their values if needed.
		These passes through the list are repeated until no swaps had to be performed during a pass, meaning that the list has become fully sorted.
		The algorithm, which is a comparison sort, is named for the way the larger elements "bubble" up to the top of the list.

		Time Complexity: - Worst Case:   O(n^2) (quadratic)
						 - Average Case: O(n^2) (quadratic)
						 - Best Case:    O(n)   (linear) {Only if it's optimized for the best case scenario}

		Space Complexity: - All Cases:   O(1)   (constant) {In place}
	*/

	for i := 0; i < len(array); i++ { // Access each element
		swapped := false
		for j := 0; j < len(array)-i-1; j++ { // Compare adjacent elements
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j] // Swap the values
				swapped = true                              // We swapped a element
			}
		}
		if !swapped { // Improve the best case time complexity
			break // No swaps were performed, therefore the array is sorted, exit the loop
		}
	}
}
