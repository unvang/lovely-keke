package week08

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

//位1的个数;hamming重量
//O(k)时间复杂度，k是uint32位数
func hammingWeight(num uint32) int {
	res :=0
    for i:=0 ;i<32;i++{
		if 1<<i & num >0{
			res++
		}
	}
	return res
}
//O(logn)时间复杂度，比如7，111，只需查询三次，n = n & (n-1) 刚好把二进制末位1变成0
func hammingWeight(num uint32) int {
	res :=0
	for ;num>0;num &= (num-1){
		res++
	}
	return res
}

//2的幂,n是2的幂，n的二进制里面只有一个1
func isPowerOfTwo(n int) bool {
	return n>0 && n & (n-1) ==0
}
func isPowerOfTwo(n int) bool {
	return n>0 && n &-n ==n
}
func isPowerOfTwo(n int) bool {
	big := 1 << 30
	return n > 0 && big%n == 0
}