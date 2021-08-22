package week03

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var idxMap map[int]int
var preorderIdx int

func buildTree(preorder []int, inorder []int) *TreeNode {
	idxMap = make(map[int]int)
	preorderIdx = 0
	for i, v := range inorder {
		idxMap[v] = i
	}
	return build(0, len(inorder)-1, preorder)
}

func build(left, right int, preorder []int) *TreeNode {
	if left > right {
		return nil
	}
	root := &TreeNode{Val: preorder[preorderIdx]}
	preorderIdx++
	idx := idxMap[root.Val]
	root.Left = build(left, idx-1, preorder)
	root.Right = build(idx+1, right, preorder)
	return root
}

//solution 2
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	idx := indexOf(inorder, preorder[0])
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}
}

func indexOf(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}
