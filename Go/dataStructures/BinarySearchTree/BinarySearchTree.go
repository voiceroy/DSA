package binarysearchtree

type TreeNode struct {
	Value  int
	Parent *TreeNode
	Left   *TreeNode
	Right  *TreeNode
}

type BST struct {
	Root *TreeNode
}

func MakeNode(value int) *TreeNode {
	return &TreeNode{value, nil, nil, nil}
}

func MakeBST() *BST {
	return &BST{nil}
}

func (bstree *BST) Depth(node *TreeNode) int {
	var currentNode = bstree.Root
	var currentDepth = 0

	for currentNode != nil {
		if currentNode.Value == node.Value {
			return currentDepth
		} else if node.Value < currentNode.Value {
			currentNode = currentNode.Left
		} else if node.Value > currentNode.Value {
			currentNode = currentNode.Right
		}
		currentDepth++
	}

	return -1
}

func (bstree *BST) Height(node *TreeNode) int {
	if node == nil {
		return -1
	}

	var leftHeight = bstree.Height(node.Left)
	var rightHeight = bstree.Height(node.Right)

	return max(leftHeight, rightHeight) + 1
}

func (bstree *BST) Search(node *TreeNode) *TreeNode {
	var currentNode = bstree.Root

	for currentNode != nil {
		if node.Value == currentNode.Value {
			return currentNode
		} else if node.Value < currentNode.Value {
			currentNode = currentNode.Left
		} else if node.Value > currentNode.Value {
			currentNode = currentNode.Right
		}
	}

	return nil
}

func (bstree *BST) Maximum(node *TreeNode) *TreeNode {
	for node.Right != nil {
		node = node.Right
	}

	return node
}

func (bstree *BST) Minimum(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

func (bstree *BST) Successor(node *TreeNode) *TreeNode {
	if node.Right != nil {
		return bstree.Minimum(node.Right)
	}

	var successor = node.Parent
	for successor != nil && node == successor.Right {
		node = successor
		successor = successor.Parent
	}

	return successor
}

func (bstree *BST) Predecessor(node *TreeNode) *TreeNode {
	if node.Left != nil {
		return bstree.Maximum(node.Left)
	}

	var predecessor = node.Parent
	for predecessor != nil && node == predecessor.Left {
		node = predecessor
		predecessor = predecessor.Parent
	}

	return predecessor
}

func (bstree *BST) Insert(node *TreeNode) *TreeNode {
	if bstree.Root == nil {
		bstree.Root = node
		return bstree.Root
	} else {
		var currentNode = bstree.Root

		for currentNode != nil {
			if node.Value < currentNode.Value {
				if currentNode.Left != nil {
					currentNode = currentNode.Left
				} else {
					currentNode.Left = node
					currentNode.Left.Parent = currentNode
					break
				}
			} else {
				if currentNode.Right != nil {
					currentNode = currentNode.Right
				} else {
					currentNode.Right = node
					currentNode.Right.Parent = currentNode
					break
				}
			}
		}

		return node
	}
}

func (bstree *BST) shiftNodes(replace *TreeNode, replaceWith *TreeNode) {
	if replace.Parent == nil {
		bstree.Root = replaceWith
	} else if replace.Parent.Left == replace {
		replace.Parent.Left = replaceWith
	} else {
		replace.Parent.Right = replaceWith
	}

	if replaceWith != nil {
		replaceWith.Parent = replace.Parent
	}
}

func (bstree *BST) Delete(node *TreeNode) *TreeNode {
	var currentNode = bstree.Search(node)

	if currentNode == nil {
		return node
	} else {
		if currentNode.Left == nil {
			bstree.shiftNodes(currentNode, currentNode.Right)
		} else if currentNode.Right == nil {
			bstree.shiftNodes(currentNode, currentNode.Left)
		} else {
			var successor = bstree.Successor(currentNode)
			if successor.Parent != currentNode {
				bstree.shiftNodes(successor, successor.Right)
				successor.Right = currentNode.Right
				successor.Right.Parent = successor
			}

			bstree.shiftNodes(currentNode, successor)
			successor.Left = currentNode.Left
			successor.Left.Parent = successor
		}

		return currentNode
	}
}

func (bstree *BST) InOrderTraversal(node *TreeNode) []int {
	var order []int
	if node != nil {
		order = append(order, bstree.InOrderTraversal(node.Left)...)
		order = append(order, node.Value)
		order = append(order, bstree.InOrderTraversal(node.Right)...)
	}

	return order
}

func (bstree *BST) PreOrderTraversal(node *TreeNode) []int {
	var order []int
	if node != nil {
		order = append(order, node.Value)
		order = append(order, bstree.PreOrderTraversal(node.Left)...)
		order = append(order, bstree.PreOrderTraversal(node.Right)...)
	}

	return order
}

func (bstree *BST) PostOrderTraversal(node *TreeNode) []int {
	var order []int
	if node != nil {
		order = append(order, bstree.PostOrderTraversal(node.Left)...)
		order = append(order, bstree.PostOrderTraversal(node.Right)...)
		order = append(order, node.Value)
	}

	return order
}
