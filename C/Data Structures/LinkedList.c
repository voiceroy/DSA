#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>

/*
   In computer science, a linked list is a linear collection of data elements
   whose order is not given by their physical placement in memory. Instead, each
   element points to the next. It is a data structure consisting of a collection
   of nodes which together represent a sequence. In its most basic form, each
   node contains data, and a reference (in other words, a link) to the next node
   in the sequence. This structure allows for efficient insertion or removal of
   elements from any position in the sequence during iteration. More complex
   variants add additional links, allowing more efficient insertion or removal
   of nodes at arbitrary positions. A drawback of linked lists is that data
   access time is a linear function of the number of nodes for each linked list
   i.e. the access time linearly increases as nodes are added to a linked list.
   Because nodes are serially linked so a node needs to be accessed first to
   access the next node (so difficult to pipeline). Faster access, such as
   random access, is not feasible.

   Time Complexity: - Get:            O(n) (linear)
                    - InsertHead:     O(1) (constant)
                    - InsertBetween:  O(n) (linear)
                    - InsertTail:     O(1) (constant)
                    - Remove:         O(n) (linear)
                    - GetValues:      O(n) (linear)

   Space Complexity: - All Cases:     O(n) (linear)
*/

struct Node {
  int value;
  struct Node *next;
};

struct LinkedList {
  struct Node *head, *tail;
  size_t length;
};

struct Node *makeNode(int value) {
  struct Node *newNode = (struct Node *)malloc(sizeof(struct Node));

  if (newNode == NULL) {
    printf("Memory allocation failed.\n");
    exit(1);
  }

  newNode->value = value;
  newNode->next = NULL;
  return newNode;
}

struct LinkedList *makeLinkedList() {
  struct LinkedList *ll =
      (struct LinkedList *)malloc(sizeof(struct LinkedList));

  if (ll == NULL) {
    printf("Memory allocation failed.\n");
    exit(1);
  }

  ll->head = NULL;
  ll->tail = NULL;
  ll->length = 0;

  return ll;
}

int len(struct LinkedList *ll) { return ll->length; }

int get(size_t index, struct LinkedList *ll) {
  if (index < 0 || index >= ll->length) {
    printf("Out of bounds\n");
    return -1;
  }

  struct Node *currentNode = ll->head;

  for (int i = 0; i < index; i++) {
    currentNode = currentNode->next;
  }

  return currentNode->value;
}

void insertHead(int value, struct LinkedList *ll) {
  struct Node *newNode = makeNode(value);

  if (ll->head == NULL) {
    ll->head = newNode;
    ll->tail = ll->head;
  } else {
    newNode->next = ll->head;
    ll->head = newNode;
  }
  ll->length++;
}

void insertBetween(size_t index, int value, struct LinkedList *ll) {
  if (index < 0 || index >= ll->length) {
    printf("Out of bounds\n");
    return;
  }

  struct Node *currentNode = ll->head;
  struct Node *newNode = makeNode(value);

  for (int i = 0; i < index - 1; i++) {
    currentNode = currentNode->next;
  }

  newNode->next = currentNode->next;
  currentNode->next = newNode;
  ll->length++;
}

void insertTail(int value, struct LinkedList *ll) {
  struct Node *newNode = makeNode(value);

  if (ll->tail == NULL) {
    ll->head = newNode;
    ll->tail = ll->head;
  } else {
    ll->tail->next = newNode;
    ll->tail = newNode;
  }

  ll->length++;
}

int Remove(size_t index, struct LinkedList *ll) {
  if (index < 0 || index >= ll->length) {
    printf("Out of bounds\n");
    return 0;
  }
  struct Node *currentNode = ll->head;

  // Special case of removal at head
  if (index == 0) {
    if (currentNode != NULL) {
      ll->head = currentNode->next;
    } else {
      ll->head = NULL;
      ll->tail = ll->head;
    }
    ll->length--;
    int value = currentNode->value;
    return value;
  }

  // Go to the last element before the element to be removed
  for (int i = 0; i < index - 1; i++) {
    currentNode = currentNode->next;
  }

  struct Node *node = currentNode->next;
  currentNode->next = node->next;
  ll->length--;

  // Special case of removal at tail
  if (currentNode->next == NULL) {
    ll->tail = currentNode;
  }
  int value = node->value;
  free(node);

  return value;
}

int *getValues(struct LinkedList *ll) {
  struct Node *currentNode = ll->head;

  int *array = (int *)malloc(sizeof(int) * ll->length);

  for (int i = 0; i < ll->length; i++) {
    array[i] = currentNode->value;
    currentNode = currentNode->next;
  }

  return array;
}
