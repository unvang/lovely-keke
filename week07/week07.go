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

var res [][]string

func solveNQueens(n int) [][]string {
	col := make(map[int]bool, 0)
	dia1 := make(map[int]bool, 0)
	dia2 := make(map[int]bool, 0)
	res = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	solve(n, 0, queens, col, dia1, dia2)
	return res
}

func solve(n, row int, queens []int, col, dia1, dia2 map[int]bool) {
	if row == n {
		str := boardPrint(queens, n)
		res = append(res, str)
		return
	}
	for i := 0; i < n; i++ {
		if col[i] {
			continue
		}
		diagonal1 := row - i
		if dia1[diagonal1] {
			continue
		}
		diagonal2 := row + i
		if dia2[diagonal2] {
			continue
		}
		queens[row] = i
		col[i] = true
		dia1[diagonal1], dia2[diagonal2] = true, true
		solve(n, row+1, queens, col, dia1, dia2)
		queens[row] = -1
		delete(col, i)
		delete(dia1, diagonal1)
		delete(dia2, diagonal2)
	}
}

func boardPrint(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		//fmt.Println(string(row))
		board = append(board, string(row))
	}
	return board
}
