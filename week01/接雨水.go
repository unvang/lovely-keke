package week01

//暴力求解,按列求
//计算每列的高度，先求出每列左右两边最高的值，其中min减去该列本身
//就是该列的雨水值
//碰到第一列和最后一列跳过
//如果计算的结果 < 0，那么就是该列上接不到雨水
func trap(height []int) int {
	n := len(height)
	result := 0
	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			continue
		}
		rHeight, lHeight := height[i], height[i]
		for r := i + 1; r < n; r++ {
			if height[r] > rHeight {
				rHeight = height[r]
			}
		}
		for l := i - 1; l >= 0; l-- {
			if height[l] > lHeight {
				lHeight = height[l]
			}
		}
		h := min(rHeight, lHeight) - height[i]
		if h > 0 {
			result += h
		}
	}
	return result
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//双指针夹逼，按列求
//bucket是桶高，
//确定遍历这一步是：从左往右还是从右往左；hi<hj就是从左往右,桶壁位于i左侧，此时从左往右遍历，只需计算桶壁减去遍历到的值，得到列上的雨水高度
func trap2(height []int) int {
	n := len(height)
	result := 0
	i := 0
	j := n - 1
	bucket := 0
	for i <= j {
		//确定桶高
		block := min(height[i], height[j])
		if block > bucket {
			bucket = block
		}
		//从左到右还是从右到左
		if height[i] < height[j] {
			result += bucket - height[i]
			i++
		} else {
			result += bucket - height[j]
			j--
		}
	}
	return result
}

//单调栈，按行求
//单调递减的入栈，碰到一个比栈顶大的，栈顶就可以出栈结算雨水了
//1.底部：cur;结算的地面底部就是当前栈顶
//2.左右边界：r,l;左右两边小的那个是桶壁，左边就是当前栈顶弹出后的栈顶，右边就是当前遍历到的 i；得到高度 min(r,l)-cur
//3.宽度：r-l -1
func trap3(height []int) int {
	result := 0
	st := stack.NewStack()
	for i := range height {
		for st.Length != 0 && height[st.Top()] < height[i] {
			cur := st.Top()
			st.Pop()
			if st.Length == 0 {
				break
			}
			r := i
			l := st.Top()
			result += (r - l - 1) * (min(height[l], height[r]) - height[cur])
		}
		st.Push(i)
	}
	return result
}
