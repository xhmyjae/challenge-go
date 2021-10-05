package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		left := root.Left
		right := root.Right
		var queue []*TreeNode
		queue = append(queue, root)
		for i := BTreeLevelCount(root); i > 0; i-- {
			// for BTreeLevelCount(root) > 0 {
			curr := queue[0]
			queue = queue[1:]
			// for i := BTreeLevelCount(root); i > 0; i-- {
			if left != nil {
				queue = append(queue, left)
				f(curr)
				// BTreeApplyByLevel(left, f)
			}
			if right != nil {
				queue = append(queue, right)
				f(curr)
				// BTreeApplyByLevel(right, f)
			}
			// }
		}
	}
}
