"""
    In computer science, a binary search tree (BST), also called an ordered or sorted binary tree,
    is a rooted binary tree data structure with the key of each internal node being greater than all
    the keys in the respective node's left subtree and less than the ones in its right subtree.
    The time complexity of operations on the binary search tree is linear with respect to the height of the tree.
    The performance of a binary search tree is dependent on the order of insertion of the nodes into the tree since
    arbitrary insertions may lead to degeneracy; several variations of the binary search tree can be built with guaranteed worst-case performance.
    The basic operations include: search, traversal, insert and delete. BSTs with guaranteed worst-case complexities perform better than an unsorted array,
    which would require linear search time.

    Time Complexity: - Worst Case {Degenerate Tree}:
                         - Search: O(n)      (linear)
                         - Insert: O(n)      (linear)
                         - Delete: O(n)      (linear)

                     - Average Case {Non-Degenerate Tree}:
                         - Search: O(log(n)) (logarithmic)
                         - Insert: O(log(n)) (logarithmic)
                         - Delete: O(log(n)) (logarithmic)

    Space Complexity: - All Cases: O(n)      (linear)
"""


class Node:
    def __init__(self, value, parent=None):
        self.value = value
        self.parent = parent
        self.left = None
        self.right = None


class BST:
    def __init__(self):
        self.root = None

    def depth(self, node: Node) -> int:
        """
        The depth of a node is the number of edges present in path from the root node of a tree to that node
        """
        currentNode = self.root
        currentDepth = 0

        while currentNode:
            # Found the node
            if node.value == currentNode.value:
                return currentDepth
            # Go left
            elif node.value <= currentNode.value:
                currentNode = currentNode.left
            # Go right
            else:
                currentNode = currentNode.right
            # The node is further down
            currentDepth += 1

        # Node doesn't exist
        return -1

    def height(self, node: Node) -> int:
        """
        The height of a node in a binary tree is the largest number of edges in a path from a leaf node to a target node.
        """
        # The parent is a leaf node, so it has 0 height
        # So to counter the +1 in the general case, return -1 in the base case
        if node is None:
            return -1
        else:
            # Get the height of left subtree
            leftHeight = self.height(node.left)
            # Get the height of right subtree
            rightHeight = self.height(node.right)

            # Return the maximum height
            return max(leftHeight, rightHeight) + 1

    def search(self, node: Node) -> Node:
        """
        Search for a node with the value `key`.
        """
        if node is None:
            return None

        currentNode = self.root

        # Search until we reach a leaf node
        while currentNode:
            # Found the node
            if node.value == currentNode.value:
                return currentNode
            # Node is in the left subtree
            elif node.value <= currentNode.value:
                currentNode = currentNode.left
            # Node is in the right subtree
            else:
                currentNode = currentNode.right

        # Node not found
        return None

    def maximum(self, node: Node) -> Node:
        """Find the right most node in the node's right subtree"""

        while node.right:
            node = node.right

        return node

    def minimum(self, node: Node) -> Node:
        """Find the left most node in the node's left subtree"""

        while node.left:
            node = node.left

        return node

    def successor(self, node: Node) -> Node:
        """
        Successor of a node is the node with the smallest key that is greater than the key of the given node.
        It's the node you would encounter just after the given node when performing an in-order traversal of the BST.
        """
        # Return the leftmost node in the node's right subtree if it exists
        if node.right:
            return self.minimum(node.right)

        # No right subtree, find the maximum value in the ancestors
        successor = node.parent
        # Go up until we are not the left child
        # If the node is not the right child of its parent, then the parent is its successor
        while successor and node == successor.right:
            node = successor
            successor = node.parent

        return successor

    def predecessor(self, node: Node) -> Node:
        """
        Predecessor of a node is the node with the largest key that is smaller than the key of the given node.
        It's the node you would encounter just before the given node when performing an in-order traversal of the BST.
        """
        # Return the rightmost node in the node's left subtree if it exists
        if node.left:
            return self.maximum(node.left)

        # No left subtree, find the minimum value in the ancestors
        predecessor = node.parent
        # Go up until we are not the right child
        # If the node is not the left child of its parent, then the parent is its predecessor
        while predecessor and node == predecessor.left:
            node = predecessor
            predecessor = node.parent

        return predecessor

    def insert(self, value) -> Node:
        """
        Inserts a node in the tree with the value `value` and return the inserted node.
        """
        # The tree is empty, insert a root node with `value`
        if self.root is None:
            self.root = Node(value)
        else:
            currentNode = self.root

            while currentNode:
                # Go left
                if value <= currentNode.value:
                    # Current node has a left child, go down
                    if currentNode.left is not None:
                        currentNode = currentNode.left
                    # Current node has no left child, insert here
                    else:
                        currentNode.left = Node(value, currentNode)
                        return currentNode.left
                # Go right
                else:
                    # Current node has a right child, go down
                    if currentNode.right is not None:
                        currentNode = currentNode.right
                    # Current node has no right child, insert here
                    else:
                        currentNode.right = Node(value, currentNode)
                        return currentNode.right

    def _shiftNodes(self, replace: Node, replaceWith: Node):
        """
        Helper function for deleting nodes.
        Only shifts the replaceWith to replace's place. Nothing else
        """
        # We are shifting the root, the only node with no parent
        if replace.parent is None:
            self.root = replaceWith
        # The node to be replaced is a left child
        elif replace == replace.parent.left:
            replace.parent.left = replaceWith
        # The node to be replaced is a right child
        else:
            replace.parent.right = replaceWith

        # The replaceWith node has a value, and so must have a parent
        if replaceWith is not None:
            replaceWith.parent = replace.parent

    def delete(self, node: Node) -> Node:
        """
        Delete a node with the value `value`
        """
        currentNode = self.search(node)

        # There's no node with value `value`
        if currentNode is None:
            return currentNode

        # currentNode has no left children, replace it with its right child
        if currentNode.left is None:
            self._shiftNodes(currentNode, currentNode.right)
        # currentNode has no right children, replace it with its left child
        elif currentNode.right is None:
            self._shiftNodes(currentNode, currentNode.left)
        # currentNode has either no children, or has both children
        else:
            # Replace currentNode with its successor
            successor = self.successor(currentNode)

            # If the successor is not the currentNode's immediate right child
            if successor.parent != currentNode:
                # Shift successor's right child in place of successor
                self._shiftNodes(successor, successor.right)
                # Assign the currentNode's right subtree to the successor
                successor.right = currentNode.right
                # Make successor the parent of the right subtree
                successor.right.parent = successor

            # Successor is the currentNode's immediate right child
            # Shift successor to the currentNode's place
            self._shiftNodes(currentNode, successor)
            # Assign the currentNode's left subtree to the successor
            successor.left = currentNode.left
            # Make successor the parent of the left subtree
            successor.left.parent = successor

        return currentNode

    def inOrderTraversal(self, node: Node) -> list:
        """
        Perform an inOrder traversal of a binary tree and return a list containing it.
        """
        order = []
        if node is not None:
            # inOrderTraverse the left subtree
            order.extend(self.inOrderTraversal(node.left))
            # Append the current node
            order.append(node.value)
            # inOrderTraverse the right subtree
            order.extend(self.inOrderTraversal(node.right))

        return order

    def preOrderTraversal(self, node: Node) -> list:
        """
        Perform an preOrder traversal of a binary tree and return a list containing it.
        """
        order = []
        if node is not None:
            # Append the current node
            order.append(node.value)
            # preOrderTraverse the left subtree
            order.extend(self.preOrderTraversal(node.left))
            # preOrderTraverse the right subtree
            order.extend(self.preOrderTraversal(node.right))

        return order

    def postOrderTraversal(self, node: Node) -> list:
        """
        Perform an postOrder traversal of a binary tree and return a list containing it.
        """
        order = []
        if node is not None:
            # postOrderTraverse the left subtree
            order.extend(self.postOrderTraversal(node.left))
            # postOrderTraverse the right subtree
            order.extend(self.postOrderTraversal(node.right))
            # Append the current node
            order.append(node.value)

        return order
