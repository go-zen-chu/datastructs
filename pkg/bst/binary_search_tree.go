package bst

import (
	"fmt"
	"strings"
)

// BinarySearchTree (BST) is a binary search tree struct that has a root node
type BinarySearchTree struct {
	root *TreeNode
}

type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

// NewBinarySearchTree creates BinarySearchTree
func NewBinarySearchTree(val int) *BinarySearchTree {
	return &BinarySearchTree{
		root: &TreeNode{
			Val:    val,
			Parent: nil, // there is no parent for root node
		},
	}
}

// Add adds one node to BST
func (bt *BinarySearchTree) Add(val int) {
	if bt.root == nil {
		bt.root = &TreeNode{Val: val}
	} else {
		bt.addTreeNode(bt.root, &TreeNode{Val: val})
	}
}

func (bt *BinarySearchTree) addTreeNode(tree *TreeNode, node *TreeNode) *TreeNode {
	if node.Val <= tree.Val {
		if tree.Left == nil {
			tree.Left = node
			node.Parent = tree
			return tree
		} else {
			return bt.addTreeNode(tree.Left, node)
		}
	} else {
		if tree.Right == nil {
			tree.Right = node
			node.Parent = tree
			return tree
		} else {
			return bt.addTreeNode(tree.Right, node)
		}
	}
}

// Find finds node that are equal to given value
// Returns nil if no TreeNode found
func (bt *BinarySearchTree) Find(val int) *TreeNode {
	return bt.find(bt.root, val)
}

func (bt *BinarySearchTree) find(tree *TreeNode, val int) *TreeNode {
	if tree.Val == val {
		return tree
	} else if val < tree.Val {
		if tree.Left == nil {
			return nil
		} else {
			return bt.find(tree.Left, val)
		}
	} else {
		if tree.Right == nil {
			return nil
		} else {
			return bt.find(tree.Right, val)
		}
	}
}

// FindMax finds max value node below the TreeNode
func FindMax(tree *TreeNode) *TreeNode {
	if tree.Right != nil {
		return FindMax(tree.Right)
	} else {
		return tree
	}
}

// FindMin finds min value node below the TreeNode
func FindMin(tree *TreeNode) *TreeNode {
	if tree.Left != nil {
		return FindMin(tree.Left)
	} else {
		return tree
	}
}

// Delete delete node if it exists.
// Returns false if the given node does not exists.
func (bt *BinarySearchTree) Delete(val int) bool {
	node := bt.Find(val)
	if node == nil {
		return false
	} else {
		if node.Left != nil {
			leftMaxNode := FindMax(node.Left)
			if node.Left == leftMaxNode {
				if node.Parent == nil { // if node is root
					leftMaxNode.Parent = nil
					bt.root = leftMaxNode
				} else if node.Parent.Left != nil && node.Parent.Left.Val == node.Val {
					node.Parent.Left = leftMaxNode
					leftMaxNode.Parent = node.Parent
				} else {
					node.Parent.Right = leftMaxNode
					leftMaxNode.Parent = node.Parent
				}
				// join node.Right to lefMaxNode if exists
				if node.Right != nil {
					leftMaxNode.Right = node.Right
					node.Right.Parent = leftMaxNode
				}
			} else {
				// first, remove leftMaxNode from current position
				if leftMaxNode.Left != nil {
					leftMaxNode.Parent.Right = leftMaxNode.Left
					leftMaxNode.Left.Parent = leftMaxNode.Parent
				} else {
					leftMaxNode.Parent.Right = nil
				}
				// second, replace node with leftMaxNode
				leftMaxNode.Left = node.Left
				node.Left.Parent = leftMaxNode
				if node.Right != nil {
					leftMaxNode.Right = node.Right
					node.Right.Parent = leftMaxNode
				}
				// if root node, node.Parent is nil
				if node.Parent != nil {
					leftMaxNode.Parent = node.Parent
					if node.Parent.Left != nil && node.Parent.Left.Val == node.Val {
						node.Parent.Left = leftMaxNode
					} else {
						node.Parent.Right = leftMaxNode
					}
				} else {
					leftMaxNode.Parent = nil
					bt.root = leftMaxNode
				}
			}
		} else if node.Right != nil { // no left child node but has right child node
			rightMinNode := FindMin(node.Right)
			if node.Right == rightMinNode {
				if node.Parent == nil { // if node is root
					rightMinNode.Parent = nil
					bt.root = rightMinNode
				} else if node.Parent.Left != nil && node.Parent.Left.Val == node.Val {
					node.Parent.Left = rightMinNode
					rightMinNode.Parent = node.Parent
				} else {
					node.Parent.Right = rightMinNode
					rightMinNode.Parent = node.Parent
				}
			} else {
				if rightMinNode.Right != nil {
					rightMinNode.Parent.Left = rightMinNode.Right
					rightMinNode.Right.Parent = rightMinNode.Parent
				} else {
					rightMinNode.Parent.Left = nil
				}
				rightMinNode.Right = node.Right
				node.Right.Parent = rightMinNode
				if node.Left != nil {
					rightMinNode.Left = node.Left
					node.Left.Parent = rightMinNode
				}
				if node.Parent != nil {
					rightMinNode.Parent = node.Parent
					if node.Parent.Right != nil && node.Parent.Right.Val == node.Val {
						node.Parent.Right = rightMinNode
					} else {
						node.Parent.Left = rightMinNode
					}
				} else {
					rightMinNode.Parent = nil
					bt.root = rightMinNode
				}
			}
		} else { // no child
			if node.Parent == nil {
				bt.root = nil
			} else {
				if node.Parent.Left != nil && node.Parent.Left.Val == node.Val {
					node.Parent.Left = nil
				} else {
					node.Parent.Right = nil
				}
				node.Parent = nil
			}
		}
		return true
	}
}

// Equal compares other BST and check if it's same
func (bt *BinarySearchTree) Equal(target *BinarySearchTree) bool {
	return equal(bt.root, target.root)
}

func equal(srcNode *TreeNode, dstNode *TreeNode) bool {
	if srcNode.Val != dstNode.Val {
		return false
	}
	leftEqual := false
	if srcNode.Left != nil {
		if dstNode.Left == nil {
			return false
		} else {
			leftEqual = equal(srcNode.Left, dstNode.Left)
		}
	} else {
		if dstNode.Left != nil {
			return false
		} else {
			leftEqual = true
		}
	}
	rightEqual := false
	if srcNode.Right != nil {
		if dstNode.Right == nil {
			return false
		} else {
			rightEqual = equal(srcNode.Right, dstNode.Right)
		}
	} else {
		if dstNode.Right != nil {
			return false
		} else {
			rightEqual = true
		}
	}
	if leftEqual && rightEqual {
		return true
	}
	return false
}

// Clone clones BST recursively
func (bt *BinarySearchTree) Clone() *BinarySearchTree {
	t := &TreeNode{Val: bt.root.Val}
	clone(bt.root, t)
	return &BinarySearchTree{root: t}
}

func clone(srcTree *TreeNode, dstTree *TreeNode) {
	if srcTree.Left != nil {
		dstTree.Left = &TreeNode{Val: srcTree.Left.Val, Parent: dstTree}
		clone(srcTree.Left, dstTree.Left)
	}
	if srcTree.Right != nil {
		dstTree.Right = &TreeNode{Val: srcTree.Right.Val, Parent: dstTree}
		clone(srcTree.Right, dstTree.Right)
	}
}

// String method prints tree for visualization. Implements Stringer interface
func (bt *BinarySearchTree) String() string {
	if bt.root == nil {
		return "no root"
	}
	b := &strings.Builder{}
	var err error
	if b, err = stringTree(b, 0, bt.root); err != nil {
		panic(err)
	}
	return b.String()
}

// stringTree visualize tree with depth-first search
func stringTree(builder *strings.Builder, depth int, node *TreeNode) (*strings.Builder, error) {
	var err error
	for i := 0; i < depth; i++ {
		if _, err = builder.WriteString(" "); err != nil {
			return builder, err
		}
	}
	fmt.Fprintf(builder, "└ %d\n", node.Val)

	if node.Left != nil {
		if builder, err = stringTree(builder, depth+1, node.Left); err != nil {
			return builder, err
		}
	}
	if node.Right != nil {
		if builder, err = stringTree(builder, depth+1, node.Right); err != nil {
			return builder, err
		}
	}
	return builder, nil
}
