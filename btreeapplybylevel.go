package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		var queue []*TreeNode
		queue = append(queue, root)
		for i := BTreeLevelCount(root); i > 0; i-- {
			// for i := len(queue); i > 0; i-- {
			// for len(queue) > 0 {
			curr := queue[0]
			f(curr)
			// for i := BTreeLevelCount(root); i > 0; i-- {
			if curr.Left != nil {
				queue = append(queue, curr.Left)
				// BTreeApplyByLevel(left, f)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
				// BTreeApplyByLevel(right, f)
			}
			RemoveElmt(queue, 0)
			// }
		}
	}
}

func RemoveElmt(arr []*TreeNode, index int) {
	arr = append(arr[:index], arr[index+1:]...)
}
