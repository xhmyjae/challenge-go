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

import (
	"github.com/01-edu/z01"
)

func BTreeApplyByLevel(root *TreeNode, f interface{}) {

	novo := BTreeLevelCount(root)
	for d := 1; d <= novo; d++ {

		PrintNodesLevel(root, d, f)
	}
}
func PrintNodesLevel(root *TreeNode, level int, f interface{}) {
	if root != nil {
		if level == 1 {
			ar := []interface{}{root.Data}
			z01.Call(f, ar)

		} else if level > 1 {
			PrintNodesLevel(root.Left, level-1, f)
			PrintNodesLevel(root.Right, level-1, f)
		}
	}
}
