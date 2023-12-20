package algorithms

/*
	In computer science, heapsort is a comparison-based sorting algorithm which can be thought of as "an implementation of selection sort using the right data structure."
	Like selection sort, heapsort divides its input into a sorted and an unsorted region, and it iteratively shrinks the unsorted region by extracting the largest element from it and inserting it into the sorted region.
	Unlike selection sort, heapsort does not waste time with a linear-time scan of the unsorted region; rather, heap sort maintains the unsorted region in a heap data structure to efficiently find the largest element in each step.

	   Time Complexity: - Worst Case:      O(nlogn) (linear logarithmic)
	                    - Average Case:    O(nlogn) (linear logarithmic)
	                    - Best Case:       O(nlogn) (linear logarithmic)

	   Space Complexity: - All Cases:      O(1)     (constant)
*/

func HeapSort(array []int) {
	makeMaxHeap(array)

	for i := len(array) - 1; i > 0; i-- {
		swap(0, i, array)
		heapify(array, 0, i)
	}
}

func makeMaxHeap(array []int) {
	for i := len(array)/2 - 1; i >= 0; i-- {
		heapify(array, i, len(array))
	}
}

func swap(a, b int, array []int) {
	array[a], array[b] = array[b], array[a]
}

func heapify(array []int, index, heapSize int) {
	leftChild := 2*index + 1
	rightChild := leftChild + 1
	largestValueIndex := index

	if heapSize > leftChild && array[largestValueIndex] < array[leftChild] {
		largestValueIndex = leftChild
	}

	if heapSize > rightChild && array[largestValueIndex] < array[rightChild] {
		largestValueIndex = rightChild
	}

	if largestValueIndex != index {
		swap(index, largestValueIndex, array)
		heapify(array, largestValueIndex, heapSize)
	}
}
