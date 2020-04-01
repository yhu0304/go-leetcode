package invert_binary_tree_226

import . "go-leetcode/structures"

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
