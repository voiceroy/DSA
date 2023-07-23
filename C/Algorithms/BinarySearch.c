#include <stddef.h>

int BinarySearch(int array[], size_t array_size, int element) {
  /*
    Binary search works on sorted arrays. Binary search begins by
    comparing an element in the middle of the array with the target value.

    If the target value matches the element, its position in the array is
    returned. If the target value is less than the element, the search
    continues in the lower half of the array.

    If the target value is greater than the element, the search continues in the
    upper half of the array. By doing this, the algorithm eliminates the half in
    which the target value cannot lie in each iteration.

    Time Complexity:  - Worst Case:   O(log(n)) (logarithmic)
                      - Average Case: O(log(n)) (logarithmic)
                      - Best Case:    O(1)      (constant)

    Space Complexity: - All Cases:    O(1)      (constant) {Only calculates
    three values, no matter the array size}

    Note: This implementation doesn't handle duplicate values
  */

  int low = 0, high = array_size; // Set the boundaries

  do {
    int middle =
        low + (high - low) / 2; // Calculate the middle point every iteration

    if (array[middle] == element) {
      return middle; // We found the element
    } else if (array[middle] > element) {
      high = middle; // The element is smaller than the current value, set the
                     // high boundary to middle
    } else {
      low = middle + 1; // The element is larger than the current value, set the
                        // low boundary to middle
    }
  } while (low < high);

  return -1; // The element is not in the list
}
