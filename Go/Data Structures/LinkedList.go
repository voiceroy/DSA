package data_structures

import (
	"errors"
)

/*
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
                    - InsertHead:     O(1) (constant)
                    - InsertBetween:  O(n) (linear)
                    - InsertTail:     O(1) (constant)
                    - Remove:         O(n) (linear)
                    - GetValues:      O(n) (linear)

   Space Complexity: - All Cases:     O(n) (linear)
*/

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head   *node
	tail   *node
	length uint
}

func MakeLinkedList() *linkedList {
	// Instantiate a empty linked list
	return &linkedList{nil, nil, 0}
}

func (ll *linkedList) Len() uint {
	// Returns the length of the linked list
	return ll.length
}

func (ll *linkedList) Get(index uint) (int, error) {
	// Returns a value at index i if exists or None
	if index < 0 || index >= ll.length {
		return 0, errors.New("Out of bounds")
	}
	var currentNode = ll.head

	for i := uint(0); i < index; i++ {
		currentNode = currentNode.next
	}

	return currentNode.value, nil
}

func (ll *linkedList) InsertHead(value int) {
	// Insert at the beginning of the linked list
	var newNode = node{value, nil}

	if ll.head == nil {
		ll.head = &newNode
		ll.tail = ll.head
	} else {
		newNode.next = ll.head
		ll.head = &newNode
	}
	ll.length++
}

func (ll *linkedList) InsertBetween(index uint, value int) error {
	// Insert between the elements of the linked list
	if index < 0 || index >= ll.length {
		return errors.New("Out of bounds")
	}

	var currentNode = ll.head
	var newNode = node{value, nil}

	for i := uint(0); i < index-1; i++ {
		currentNode = currentNode.next
	}

	newNode.next = currentNode.next
	currentNode.next = &newNode
	ll.length++
	return nil
}

func (ll *linkedList) InsertTail(value int) {
	// Inserts at the tail of the linked list
	var newNode = node{value, nil}

	if ll.tail == nil {
		ll.head = &newNode
		ll.tail = ll.head
	} else {
		ll.tail.next = &newNode
		ll.tail = &newNode
	}
	ll.length++
}

func (ll *linkedList) Remove(index uint) (int, error) {
	// Removes the element at index i if exists else None
	if index < 0 || index >= ll.length {
		return 0, errors.New("Out of bound")
	}

	// Special case of removal at head
	if index == 0 {
		var node = ll.head

		if node != nil {
			ll.head = node.next
		} else {
			ll.head = nil
			ll.tail = ll.head
		}

		ll.length--
		return node.value, nil
	}

	var currentNode = ll.head
	// Go to the last element before the element to be removed
	for i := uint(0); i < index-1; i++ {
		currentNode = currentNode.next
	}

	var node = currentNode.next
	currentNode.next = node.next
	ll.length--

	// Special case of removal at tail
	if currentNode.next == nil {
		ll.tail = currentNode
	}

	return node.value, nil
}

func (ll *linkedList) GetValues() []int {
	// Returns the contents of the linked list in array
	var currentNode = ll.head
	var array = make([]int, 0, ll.length)

	for currentNode != nil {
		array = append(array, currentNode.value)
		currentNode = currentNode.next
	}

	return array
}
