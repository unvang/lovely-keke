package thesecondtime

import (
	"fmt"
	"sort"
)

//array
//--盛水最多的容器
//给一个数组，index是x横坐标，value是y竖坐标，算组成容器的最大值

//1.确认题目：给的参数数组：非负整数数组；边界：无需考虑边界
//2.思路：组成容器最大值通过两个值计算得到，x轴长度 * y轴高度；
//那么a：算出x,y所有组合，得到最大值，使用两层for-loop
//b:双指针，i，j代表左右边界，计算：min(arr[i],arr[j]) * (j-i)，
//长度：对于数组index,x轴必然递减，所以可以遍历一遍
//高度：由于容器值是由i,j中小的那个决定的，所以移动小的那个，高度一定会改变（不考虑相等情况
//一层for-loop，计算最大值

//双指针解法
func maxArea(height []int) int {
	n := len(height)
	res := 0

	for i, j := 0, n-1; i < j; {
		res = max(res, min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//暴力解法
func maxAreaSilly(height []int) int {
	n := len(height)
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			res = max(res, min(height[j], height[i])*(j-i))
		}
	}

	return res
}

//总结：数组 index 连续，单调；可以利用这点简化暴力解法

//--移动零
//给定一个数组 nums，将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//1.审题：限制：必须在原数组上操作，减少操作次数；边界：不考虑
//2.思路: i，j初始都在数组头，j记录下一个非零放的位置，i遍历数组
//伪代码如下：
// for i,j:=0;i<n;i++{
// 	if arr[i]!=0{
// 		if i!=j{ //处于同一位置就没必要交换了
// 			swap(arr[i],arr[j])
// 		}
// 		j++
// 	}
// }
func moveZeroes(nums []int) {
	n := len(nums)
	for i, j := 0, 0; i < n; i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}

//3.总结：关键在于让 j 记录下一个非零值可以放的位置,
//i 和 j 初始都是 0, i遍历时, val不为0时，j要往前走，为 0 时即为下一个非零放的位置，停住

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

//两数之和；做法有两次for-loop 时间复杂度n方，map n;
func twoSum(nums []int, target int) []int {
	mmap := make(map[int]int, 0)
	for i, v := range nums {
		if po, ok := mmap[target-v]; ok {
			return []int{i, po}
		} else {
			mmap[v] = i
		}
	}
	return nil
}

//此处想了一下之前双指针解法，
//排序后，前后指针夹逼，然后发现题目要求返回数组下标，排序后下标会被打乱，如果用hash保存下标映射，空间增加
func twoSumErr(nums []int, target int) []int {
	n := len(nums)
	sort.Ints(nums)
	for i, j := 0, n-1; i < j; {
		if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			return []int{i, j}
		}
	}
	return nil
}

//三数之和
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n-2; first++ {
		// 需要和上一次枚举的数不相同 + 局部优化
		if first > 0 && nums[first] == nums[first-1] || nums[first]+nums[n-1]+nums[n-2] < 0 {
			continue
		}
		target := -1 * nums[first]
		// 枚举 b
		for second, third := first+1, n-1; second < third; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] || nums[second]+nums[n-1] < target{
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans

//这种解法存在导致内存过大panic；两数之和改的，两数之和不用判重，而且返回的下标
func threeSumSilly(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	for i := 0; i < n-1; i++ {
		//局部优化，不然直接内存过大导致panic
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		mmap := make(map[int]int, 0)
		for j := i + 1; j < n; j++ {
			if po, ok := mmap[0-nums[i]-nums[j]]; ok {
				res = append(res, []int{nums[i], nums[po], nums[j]})
			} else {
				mmap[nums[j]] = j
			}
		}
	}
	tmap := make(map[string][]int, 0)
	for _, v := range res {
		sort.Ints(v)
		tmap[fmt.Sprintf("%d%d%d", v[0], v[1], v[2])] = v
	}
	fRes := [][]int{}
	for _, v := range tmap {
		fRes = append(fRes, v)
	}
	return fRes
}

//四数之和
func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		//此处可以局部优化
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			newTarget := target - nums[i] - nums[j]
			for head, tail := j+1, n-1; head < tail; head++ {
				if head > j+1 && nums[head] == nums[head-1] {
					continue
				}
				for head < tail && nums[head]+nums[tail] > newTarget {
					tail--
				}
				if head == tail {
					break
				}
				if nums[head]+nums[tail] == newTarget {
					res = append(res, []int{nums[i], nums[j], nums[head], nums[tail]})
				}
			}
		}
	}
	return res
}

//四数之和局部优化
func fourSumBetter(nums []int, target int) [][]int {
	res := [][]int{}
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			newTarget := target - nums[i] - nums[j]
			for head, tail := j+1, n-1; head < tail; head++ {
				if head > j+1 && nums[head] == nums[head-1] {
					continue
				}
				for head < tail && nums[head]+nums[tail] > newTarget {
					tail--
				}
				if head == tail {
					break
				}
				if nums[head]+nums[tail] == newTarget {
					res = append(res, []int{nums[i], nums[j], nums[head], nums[tail]})
				}
			}
		}
	}
	return res
}

//反转链表
//有意思的递归，举例1-2-3-4-5，ret一直是5，返回后head会从 4 到 1，
//第一次ret返回 5 后，head是4，让4的next 5，指向 4，刚好把 5 的指针反转指向前
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	ret := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}
//递归
//此处有个点需要特别注意
//使用 var prev *ListNode prev是nil，使用 prev:=new(ListNode),prev会初始化，不是nil,val是0，
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}


//两两交换链表中的节点
//递归，典型的模版
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	one := head
	two := head.Next
	three := two.Next
	two.Next = one
	one.Next = swapPairs(three)
	return two
}
//迭代
//最好看官方图解
func swapPairs(head *ListNode) *ListNode {
	dummyNode := &ListNode{0,head}
	temp := dummyNode //返回dummyNode.Next 保持不变，temp遍历
	for temp.Next != nil && temp.Next.Next != nil{ //终止条件
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1 //node1已经替换node2位置，下一组，从原node2，现node1位置开始
	}
	return dummyNode.Next
}