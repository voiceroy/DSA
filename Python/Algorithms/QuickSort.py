def QuickSort(array: list) -> list:
    """
    Quicksort is a divide-and-conquer algorithm.
    It works by selecting a 'pivot' element from the array and partitioning the other elements into two sub-arrays, according to whether they are less than or greater than the pivot.
    For this reason, it is sometimes called partition-exchange sort. The sub-arrays are then sorted recursively.
    This can be done in-place, requiring small additional amounts of memory to perform the sorting.

    Time Complexity: - Worst Case:      O(n^2)   (quadratic)
                     - Average Case:    O(nlogn) (linear logarithmic)
                     - Best Case:       O(nlogn) (linear logarithmic)

    Space Complexity: - All Cases:      O(n)     (linear)
    """

    # Base Case. Array of 1 or 0 elements is sorted
    if len(array) <= 1:
        return array

    # Naive implementation of the algorithm. Not implementing the other optimized versions.
    pivot = array.pop()

    lessThan = [x for x in array if x <= pivot]
    greaterThan = [x for x in array if x > pivot]

    # General Case
    lessThan = QuickSort(lessThan)
    greaterThan = QuickSort(greaterThan)

    lessThan.append(pivot)
    lessThan.extend(greaterThan)
    return lessThan
