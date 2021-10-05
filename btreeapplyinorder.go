package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		f(root.Data)
		BTreeApplyInorder(root.Left, f)
		BTreeApplyInorder(root.Right, f)
	}
}


If the current node is empty then return.
Execute the following three operations in a certain order:[5]
N: Visit the current node.
L: Recursively traverse the current node's left subtree.
R: Recursively traverse the current node's right subtree.