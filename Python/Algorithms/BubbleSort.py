def BubbleSort(array: list):
    """
    Bubble sort, is a simple sorting algorithm that repeatedly steps through the input list element by element, comparing the current element with the one after it, swapping their values if needed.
    These passes through the list are repeated until no swaps had to be performed during a pass, meaning that the list has become fully sorted.
    The algorithm, which is a comparison sort, is named for the way the larger elements "bubble" up to the top of the list.

    Time Complexity: - Worst Case:   O(n^2) (quadratic)
                                 - Average Case: O(n^2) (quadratic)
                                 - Best Case:    O(n)   (linear) {Only if it's optimized for the best case scenario}


    Space Complexity: - All Cases:   O(1)   (constant) {In place}
    """

    for i in range(
        0, len(array)
    ):  # For each iteration the last i'th element will be in its correct place.
        swapped = False  # Optimise for best case
        for j in range(
            0, len(array) - 1 - i
        ):  # After each iteration last i elements will be sorted.
            if array[j] > array[j + 1]:  # Swap the elements
                array[j], array[j + 1] = array[j + 1], array[j]
                swapped = True  # We swapped an element

        if (
            not swapped
        ):  # Best case optimisation, no elements swapped, so array is sorted.
            break
