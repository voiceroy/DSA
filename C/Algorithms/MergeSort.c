#include <stddef.h>
#include <stdlib.h>
#include <string.h>

void merge(int *sorted_array, const int *left_array, const int *right_array,
           const size_t left_size, const size_t right_size) {
  size_t l = 0;
  size_t r = 0;
  size_t i = 0;

  while (l < left_size && r < right_size) {
    if (left_array[l] >= right_array[r]) {
      sorted_array[i++] = right_array[r++];
    } else {
      sorted_array[i++] = left_array[l++];
    }
  }

  if (l < left_size) {
    memcpy(sorted_array + i, left_array + l, sizeof(int) * (left_size - l));
  }

  if (r < right_size) {
    memcpy(sorted_array + i, right_array + r, sizeof(int) * (right_size - r));
  }
}

void merge_sort(int *sorted_array, const int *array, const size_t size) {
  if (size == 1) {
    memcpy(sorted_array, array, sizeof(int));
    return;
  }

  const size_t middle = size / 2;

  int *left_sorted = malloc(sizeof(int) * middle);
  int *right_sorted = malloc(sizeof(int) * (size - middle));

  merge_sort(left_sorted, array, middle);
  merge_sort(right_sorted, array + middle, (size - middle));
  merge(sorted_array, left_sorted, right_sorted, middle, size - middle);
  free(left_sorted);
  free(right_sorted);
}
