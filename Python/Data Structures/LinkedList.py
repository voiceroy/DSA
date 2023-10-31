"""
    In computer science, a linked list is a linear collection of data elements whose order is not given by their physical placement in memory.
    Instead, each element points to the next. It is a data structure consisting of a collection of nodes which together represent a sequence.
    In its most basic form, each node contains data, and a reference (in other words, a link) to the next node in the sequence.
    This structure allows for efficient insertion or removal of elements from any position in the sequence during iteration.
    More complex variants add additional links, allowing more efficient insertion or removal of nodes at arbitrary positions.
    A drawback of linked lists is that data access time is a linear function of the number of nodes for each linked list i.e.
    the access time linearly increases as nodes are added to a linked list.
    Because nodes are serially linked so a node needs to be accessed first to access the next node (so difficult to pipeline).
    Faster access, such as random access, is not feasible.

    Time Complexity: - Get:            O(n) (linear)
                     - insertHead:     O(1) (constant)
                     - insertBetween:  O(n) (linear)
                     - insertTail:     O(1) (constant)
                     - remove:         O(n) (linear)
                     - getValues:      O(n) (linear)

    Space Complexity: - All Cases:     O(n) (linear)
"""


class Node:
    def __init__(self, value):
        """Instantiates a new linked list node with value"""
        self.value = value
        self.next = None


class LinkedList:
    def __init__(self):
        """Instantiate a empty linked list"""
        self.head = self.tail = None
        self.length = 0

    def __len__(self) -> int:
        """Returns the length of the linked list"""
        return self.length

    def get(self, i: int):
        """Returns a value at index i if exists or None"""
        # Out of bounds
        if i < 0 or i >= self.length:
            return None

        currentNode = self.head
        idx = 0

        while currentNode:
            if idx == i:
                return currentNode.value
            idx += 1
            currentNode = currentNode.next

        return None

    def insertHead(self, value):
        """Insert at the beginning of the linked list"""
        node = Node(value)

        if self.head is None:
            self.head = node
            self.tail = self.head
        else:
            node.next = self.head
            self.head = node
        self.length += 1

    def insertBetween(self, i: int, value):
        """Insert between the elements of the linked list"""
        # Out of bounds
        if i < 0 or i >= self.length:
            return None

        node = Node(value)
        currentNode = self.head
        idx = 0
        while idx < i - 1:
            currentNode = currentNode.next
            idx += 1

        node.next = currentNode.next
        currentNode.next = node

    def insertTail(self, value):
        """Inserts at the tail of the linked list"""
        node = Node(value)

        if self.tail is None:
            self.head = node
            self.tail = self.head
        else:
            self.tail.next = node
            self.tail = node
        self.length += 1

    def remove(self, i: int):
        """Removes the element at index i if exists else None"""
        # Out of bounds
        if i < 0 or i >= self.length:
            return None

        # Special case of removal at head
        if i == 0:
            node = self.head

            if node:
                self.head = node.next
            else:
                self.head = self.tail = None

            self.length -= 1
            return node.value

        currentNode = self.head
        idx = 0

        # Go to the last element before the element to be removed
        while idx < i - 1:
            currentNode = currentNode.next
            idx += 1

        node = currentNode.next
        currentNode.next = node.next
        self.length -= 1

        # Special case of removal at tail
        if i == self.length:
            self.tail = currentNode

        return node.value

    def getValues(self) -> list:
        """Returns the contents of the linked list in array"""
        result = []

        currentNode = self.head
        while currentNode:
            result.append(currentNode.value)
            currentNode = currentNode.next

        return result
