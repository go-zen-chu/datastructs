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
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewBinarySearchTree creates BinarySearchTree
func NewBinarySearchTree(val int) *BinarySearchTree {
	return &BinarySearchTree{
		root: &TreeNode{
			Val: val,
		},
	}
}

// Add adds one node to BST
func (bt *BinarySearchTree) Add(val int) {
	bt.addTreeNode(bt.root, &TreeNode{Val: val})
}

func (bt *BinarySearchTree) addTreeNode(tree *TreeNode, node *TreeNode) *TreeNode {
	if node.Val <= tree.Val {
		if tree.Left == nil {
			tree.Left = node
			return tree
		} else {
			return bt.addTreeNode(tree.Left, node)
		}
	} else {
		if tree.Right == nil {
			tree.Right = node
			return tree
		} else {
			return bt.addTreeNode(tree.Right, node)
		}
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
		dstTree.Left = &TreeNode{Val: srcTree.Left.Val}
		clone(srcTree.Left, dstTree.Left)
	}
	if srcTree.Right != nil {
		dstTree.Right = &TreeNode{Val: srcTree.Right.Val}
		clone(srcTree.Right, dstTree.Right)
	}
}

// String method prints tree for visualization. Implements Stringer interface
func (bt *BinarySearchTree) String() string {
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
