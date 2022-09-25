package utils

type TElemType int

// Definition for a binary tree node.
// Tree指向第一个TreeNode的指针
type TreeNode struct {
	Val   TElemType
	Left  *TreeNode
	Right *TreeNode
}

// Binary Search Trees(二叉搜索树)
// Methods: MakeEmpty,Find,FindMin,FindMax,Insert,Delete

// AVL Tree (Self-Balancing/Height-Balanced Binary Search Tree) (二叉平衡树)
// Methods: Height,Insert,SingleRotateWithLeft,DoubleRotateWithLeft
type AvlNode struct {
	Val    TElemType
	Left   *TreeNode
	Right  *TreeNode
	Height int
}

func (t *TreeNode) String() string {
	return "This is a tree!"
}

// B-trees
// Red Black Trees
