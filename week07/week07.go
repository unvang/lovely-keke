package week07

//又是爬楼梯

//爬楼梯
//首先写一遍递归模版
// func recursion(n int) {
// 	//terminator
// 	if level > max_level {
// 		//process result
// 		return
// 	}
// 	//process current logic
// 	process(level, param)

// 	//dirll down
// 	recursion(level+1, newpararm)

// 	//restore current status
// }
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	f1, f2, f3 := 1, 2, 0
	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3
}

//傻递归
func climbStairsSilly(n int) int {
	if n <= 2 {
		return n
	}

	return climbStairsSilly(n-1) + climbStairsSilly(n-2)
}

//加缓存
func climbStairsCache(n int) int {
	if n <= 2 {
		return n
	}
	a := make([]int, n+1)
	a[1] = 1
	a[2] = 2
	for i := 3; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	return a[n]
}

//N皇后，

//回溯模版
func backtracking(参数) {
    if (终止条件) {
        存放结果;
        return;
    }
    for (选择：本层集合中元素（树中节点孩子的数量就是集合的大小）) {
        处理节点;
        backtracking(路径，选择列表); // 递归
        回溯，撤销处理结果
    }
}

//N皇后：回溯；自顶向下编程范例

var res [][]string

//顶层函数，负责返回值初始化；矩阵slice参数初始化
func solveNQueens(n int) [][]string {
	res = [][]string{}
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	backtrack(board, 0)
	return res
}
//回溯递归函数，根据模版
func backtrack(board [][]string, row int) {
	n := len(board)
	if row == n {
		temp := make([]string, n)
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(board[i], "")
		}
		res = append(res, temp)
		return
	}
	for i := 0; i < n; i++ {
		if !isSolved(board, row, i) {
			continue
		}
		board[row][i] = "Q"
		backtrack(board, row+1)
		board[row][i] = "."
	}
}
//判断是否有效的函数；行，列，左叉，右叉
func isSolved(board [][]string, row, col int) bool {
	n := len(board)
	for i := 0; i < n; i++ {
		if board[row][i] == "Q" {
			return false
		}
	}
	for i := 0; i < n; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}