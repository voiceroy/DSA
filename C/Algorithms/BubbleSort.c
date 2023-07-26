#include <stdbool.h>
#include <stddef.h>

void BubbleSort(int array[], size_t array_size) {
  /*
  Bubble sort, is a simple sorting algorithm that repeatedly steps
  through the input list element by element, comparing the current element
  with the one after it, swapping their values if needed. These passes
  through the list are repeated until no swaps had to be performed during a
  pass, meaning that the list has become fully sorted. The algorithm, which
  is a comparison sort, is named for the way the larger elements "bubble" up
  to the top of the list.

  Time Complexity: - Worst Case:   O(n^2) (quadratic)
                   - Average Case: O(n^2) (quadratic)
                   - Best Case:    O(n)   (linear) {Only if it's optimized for
  the best case scenario}

  Space Complexity: - All Cases:   O(1)   (constant) {In place}
  */
  for (int i = 0; i < array_size;
       i++) { // Iterate n times, where in each time the last i elements are in
              // sorted order
    bool swapped = false; // Best case optimisation
    for (int j = 0; j < array_size - 1 - i;
         j++) { // Go upto N - 1 - i elements are after every iteration the last
                // i elements are in sorted order.
      if (array[j] > array[j + 1]) { // Swap the values if true
        int temp = array[j];
        array[j] = array[j + 1];
        array[j + 1] = temp;
        swapped = true;
      }
    }
    if (!swapped) { // No swaps occured, hence the array is sorted
      break;        // Break the loop
    }
  }
}
