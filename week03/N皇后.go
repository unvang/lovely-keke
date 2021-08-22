package week03

var res [][]string

//列和两条斜线是否已有，使用三个map判断
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
		col[i] = true //放在本行第几列
		dia1[diagonal1], dia2[diagonal2] = true, true
		solve(n, row+1, queens, col, dia1, dia2) //下一行

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
