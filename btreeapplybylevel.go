package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		left := root.Left
		right := root.Right
		var queue []*TreeNode
		queue = append(queue, root)
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[0:]
			// for i := BTreeLevelCount(root); i > 0; i-- {
			if left != nil {
				queue = append(queue, left)
				f(curr)
				// BTreeApplyByLevel(left, f)
			}
			if right != nil {
				queue = append(queue, right)
				// BTreeApplyByLevel(right, f)
			}
			// }
		}
	}
}
