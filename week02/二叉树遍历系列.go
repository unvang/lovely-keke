package week02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//preorder 递归
func preorderTraversal(root *TreeNode) []int {
	a := []int{}
	pre(root, &a)
	return a
}
func pre(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}
	*a = append(*a, root.Val)
	pre(root.Left, a)
	pre(root.Right, a)
}

//迭代
func preorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		st := []*TreeNode{root}
		for len(st) > 0 {
			top := st[len(st)-1]
			res = append(res, top.Val)
			//pop
			st[len(st)-1] = nil
			st = st[:len(st)-1]

			if top.Right != nil {
				st = append(st, top.Right)
			}
			if top.Left != nil {
				st = append(st, top.Left)
			}
		}
	}
	return res
}

//inOrder 递归
func inorderTraversal(root *TreeNode) []int {
	a := []int{}
	inOrder(root, &a)
	return a
}
func inOrder(node *TreeNode, a *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, a)
	*a = append(*a, node.Val)
	inOrder(node.Right, a)
}

//迭代
func inorderTraversal2(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		st := []*TreeNode{}
		cur := root
		for cur != nil || len(st) > 0 {
			if cur != nil {
				st = append(st, cur)
				cur = cur.Left
			} else {
				cur = st[len(st)-1]
				//pop
				st[len(st)-1] = nil
				st = st[:len(st)-1]

				res = append(res, cur.Val)
				cur = cur.Right
			}
		}

	}
	return res
}

//postOrder 递归
func postorderTraversal(root *TreeNode) (a []int) {
	var postOrder func(*TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		a = append(a, node.Val)
	}
	postOrder(root)
	return
}

//迭代
func postorderTraversal2(root *TreeNode) (a []int) {
	res := []int{}
	if root != nil {
		st := []*TreeNode{root}
		for len(st) > 0 {
			top := st[len(st)-1]
			res = append(res, top.Val)
			st = st[:len(st)-1]
			if top.Left != nil {
				st = append(st, top.Left)
			}
			if top.Right != nil {
				st = append(st, top.Right)
			}
		}
	}
	reverse(res)
	return res
}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//层序遍历 递归
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root != nil {
		level(&res, root, 0)
	}
	return res
}
func level(res *[][]int, node *TreeNode, l int) {
	if len(*res) <= l {
		*res = append(*res, []int{})
	}
	(*res)[l] = append((*res)[l], node.Val)
	if node.Left != nil {
		level(res, node.Left, l+1)
	}
	if node.Right != nil {
		level(res, node.Right, l+1)
	}
}

//迭代
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		res = append(res, []int{})
		for len(queue) > 0 {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val)
			queue[0] = nil
			queue = queue[1:]
		}
	}
	return res
}

//迭代使用nil标记法，统一，只需替换二行代码顺序
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	st := []*TreeNode{}
	if root != nil {
		st = append(st, root) //push
	}
	for len(st) > 0 {
		node := st[len(st)-1]
		if node != nil {
			//pop
			st[len(st)-1] = nil
			st = st[:len(st)-1]

			if node.Right != nil {
				st = append(st, node.Right)
			}
			if node.Left != nil {
				st = append(st, node.Left)
			}
			st = append(st, node)
			st = append(st, nil)
		} else {
			//pop nil
			st[len(st)-1] = nil
			st = st[:len(st)-1]

			top := st[len(st)-1]
			//pop
			st[len(st)-1] = nil
			st = st[:len(st)-1]
			res = append(res, top.Val)
		}
	}
	return res
}
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	st := []*TreeNode{}
	if root != nil {
		st = append(st, root) //push
	}
	for len(st) > 0 {
		node := st[len(st)-1]
		if node != nil {
			//pop
			st[len(st)-1] = nil
			st = st[:len(st)-1]

			if node.Right != nil {
				st = append(st, node.Right)
			}
			st = append(st, node)
			st = append(st, nil)
			if node.Left != nil {
				st = append(st, node.Left)
			}
		} else {
			//pop nil
			st[len(st)-1] = nil
			st = st[:len(st)-1]

			top := st[len(st)-1]
			//pop
			st[len(st)-1] = nil
			st = st[:len(st)-1]
			res = append(res, top.Val)
		}
	}
	return res
}

func postorderTraversal(root *TreeNode) (a []int) {
	res := []int{}
	st := []*TreeNode{}
	if root != nil {
		st = append(st, root) //push
	}
	for len(st) > 0 {
		node := st[len(st)-1]
		if node != nil {
			//pop
			st[len(st)-1] = nil
			st = st[:len(st)-1]
			st = append(st, node)
			st = append(st, nil)
			if node.Right != nil {
				st = append(st, node.Right)
			}

			if node.Left != nil {
				st = append(st, node.Left)
			}
		} else {
			//pop nil
			st[len(st)-1] = nil
			st = st[:len(st)-1]

			top := st[len(st)-1]
			//pop
			st[len(st)-1] = nil
			st = st[:len(st)-1]
			res = append(res, top.Val)
		}
	}
	return res
}
