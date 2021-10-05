package piscine

/*func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
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
*/

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	for i := 1; i <= BTreeLevelCount(root); i++ {
		AtLevel(root, i, f)
	}
}

func AtLevel(root *TreeNode, i int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if i == 1 {
		f(root.Data)
	} else if i > 1 {
		AtLevel(root.Left, i-1, f)
		AtLevel(root.Right, i-1, f)
	}
}
