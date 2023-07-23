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

	for i := 0; i < len(array); i++ { // As after each iteration only the last i items are in the correct order, to correctly sort all the N items in ascending order, loop N times
		for j := 0; j < len(array)-i-1; j++ { // Go over all the elements upto N - 1 - i, swapping any unordered adjacent elements, which results in `N-1-i`th element being in the correct order.
			swapped := false
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j] // Swap the values
				swapped = true                              // We swapped a element
			}
			if !swapped { // Improve the best case time complexity
				break // No swaps were performed, therefore the array is sorted, hence exit the loop
			}
		}
	}
}
