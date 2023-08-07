def SelectionSort(array: list):
    """
    Selection sort is an in-place comparision sorting algorithm. The algorithm divides the input list into two parts:
    a sorted sublist of items which is built up from left to right at the front (left) of the list and a sublist of the remaining unsorted items that occupy the rest of the list.
    Initially, the sorted sublist is empty and the unsorted sublist is the entire input list. The algorithm proceeds by finding the smallest (or largest, depending on sorting order) element in the unsorted sublist,
    exchanging (swapping) it with the leftmost unsorted element (putting it in sorted order), and moving the sublist boundaries one element to the right.

    Time Complexity: - Worst Case:    O(n^2) (quadratic)
                     - Average Case:  O(n^2) (quadratic)
                     - Best Case:     O(n^2) (quadratic) {Has to compare each element with all the others}

    Space Complexity: - All Cases:    O(1)   (constant) {In-place sorting algorithm}
    """

    for i in range(len(array)):  # Go through each index
        minimum = i  # Set the index of the minimum value

        for j in range(i + 1, len(array)):  # Compare with the succeeding values
            if (
                array[minimum] > array[j]
            ):  # If the current minimum element is greater than the element at index j, set the minimum element index to j
                minimum = j

        array[minimum], array[i] = (
            array[i],
            array[minimum],
        )  # Swap the values, now the array is sorted upto and including the i'th index
