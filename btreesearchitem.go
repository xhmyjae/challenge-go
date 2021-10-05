package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if elem == root.Data {
		return root
	} else {
		return nil
	}
}
