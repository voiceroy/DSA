def merge(left: list, right: list) -> list:
    result = []

    # progressively merge two sorted arrays into a single sorted array
    while len(left) and len(right):
        if left[0] >= right[0]:
            result.append(right.pop(0))
        else:
            result.append(left.pop(0))

    if len(left) == 0:
        return result + right
    return result + left


def MergeSort(array: list) -> list:
    """
    In computer science, merge sort (also commonly spelled as mergesort) is an efficient, general-purpose, and comparison-based sorting algorithm.
    Most implementations produce a stable sort, which means that the relative order of equal elements is the same in the input and output.
    Merge sort is a divide-and-conquer algorithm.

    Time Complexity: - Worst Case:		O(nlog(n)) (linear logarithmic)
                     - Average Case:	O(nlog(n)) (linear logarithmic)
                     - Best Case:		O(nlog(n)) (linear logarithmic)

    Space Complexity: - All Cases:		O(n)       (linear) {Destructive sort}
    """

    # Base case. A list of 1 or 0 elements is sorted
    if len(array) <= 1:
        return array

    length = len(array)
    middle = length // 2

    # General case. Break down the array into smaller pieces, merge sort them and finally merge them
    return merge(MergeSort(array[:middle]), MergeSort(array[middle:]))
