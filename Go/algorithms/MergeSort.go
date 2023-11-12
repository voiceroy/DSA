package algorithms

func merge(left, right []int) []int {
	var result = make([]int, 0, len(left)+len(right))

	// No leftpop in Golang, using index instead
	var l = 0
	var r = 0

	// progressively merge two sorted arrays into a single sorted array
	for len(left) > l && len(right) > r {
		if left[l] >= right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}

	if len(left) == l {
		return append(result, right[r:]...)
	}

	return append(result, left[l:]...)
}

func MergeSort(array []int) []int {
	/*
			In computer science, merge sort (also commonly spelled as mergesort) is an efficient, general-purpose, and comparison-based sorting algorithm.
			Most implementations produce a stable sort, which means that the relative order of equal elements is the same in the input and output.
		    Merge sort is a divide-and-conquer algorithm.

		    Time Complexity: - Worst Case:		O(nlog(n)) (linear logarithmic)
		                     - Average Case:	O(nlog(n)) (linear logarithmic)
		                     - Best Case:		O(nlog(n)) (linear logarithmic)


		    Space Complexity: - All Cases:		O(n)       (linear)
	*/

	// Base case. A list of 1 or 0 elements is sorted
	if len(array) <= 1 {
		return array
	}

	length := len(array)
	middle := length / 2

	// General case. Break down the array into smaller pieces, merge sort them and finally merge them
	return merge(MergeSort(array[:middle]), MergeSort(array[middle:]))
}
