#include <stddef.h>
void SelectionSort(int array[], size_t array_length) {
  /*
    Selection sort is an in-place comparision sorting algorithm. The algorithm
    divides the input list into two parts: a sorted sublist of items which is
    built up from left to right at the front (left) of the list and a sublist of
    the remaining unsorted items that occupy the rest of the list. Initially,
    the sorted sublist is empty and the unsorted sublist is the entire input
    list. The algorithm proceeds by finding the smallest (or largest, depending
    on sorting order) element in the unsorted sublist, exchanging (swapping) it
    with the leftmost unsorted element (putting it in sorted order), and moving
    the sublist boundaries one element to the right.

    Time Complexity: - Worst Case:    O(n^2) (quadratic)
                     - Average Case:  O(n^2) (quadratic)
                     - Best Case:     O(n^2) (quadratic) {Has to compare each
    element with all the others}

    Space Complexity: - All Cases:    O(1)   (constant) {In-place sorting
    algorithm}
  */

  // Go through each index
  for (int i = 0; i < array_length; i++) {
    int minimum = i; // Set the index of the minimum value

    // Compare with the succeeding values
    for (int j = i; j < array_length; j++) {

      if (array[minimum] > array[j]) {
        // If the current minimum element is greater than the
        // element at index j, set the minimum element index to j
        minimum = j;
      }
    }

    int temp = array[minimum]; // Swap the values, now the array is sorted upto
                               // and including the i'th index
    array[minimum] = array[i];
    array[i] = temp;
  }
}
