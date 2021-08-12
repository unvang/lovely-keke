package week02

type Node struct {
	Val      int
	Children []*Node
}

//level order 0ms,4.3MB
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := []*Node{root}
	for level := 0; len(queue) > 0; level++ {
		res = append(res, []int{})
		for len(queue) > 0 {
			queue = append(queue, queue[0].Children...)
			res[level] = append(res[level], queue[0].Val)
			queue[0] = nil
			queue = queue[1:]
		}
	}
	return res

}

//较慢 0ms 4.4MB
func levelOrder2(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	preQueue := []*Node{root}
	for level := 0; len(preQueue) > 0; level++ {
		res = append(res, []int{})
		curQueue := []*Node{}
		for _, node := range preQueue {
			res[level] = append(res[level], node.Val)
			curQueue = append(curQueue, node.Children...)
			preQueue = curQueue
		}
	}
	return res

}

//递归 0ms 4.2MB
func levelOrder3(root *Node) [][]int {
	res := [][]int{}
	if root != nil {
		level(root, &res, 0)
	}

	return res

}
func level(node *Node, res *[][]int, l int) {
	if len(*res) <= l {
		*res = append(*res, []int{})
	}
	(*res)[l] = append((*res)[l], node.Val)
	for _, child := range node.Children {
		level(child, res, l+1)
	}
}

//preorder 递归
func preorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	var pre func(*Node)
	pre = func(node *Node) {
		res = append(res, node.Val)
		for _, v := range node.Children {
			pre(v)
		}
	}
	pre(root)
	return res
}

//postorder 递归
func postorder(root *Node) []int {
	res := []int{}
	if root == nil {
		return res
	}
	var pre func(*Node)
	pre = func(node *Node) {
		for _, v := range node.Children {
			pre(v)
		}
		res = append(res, node.Val)
	}
	pre(root)
	return res
}
