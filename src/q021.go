package main

import (
	"fmt"
	"sort"
)

//给一组正整 数序列 N=[1,2,5,4,77,4,234,…] 任选其中3个，a,b,c， 使得a+b+c=K
//遍历A ， 问题变为 b + c = k，首位指针相加判断

func threeSum(nums []int, k int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for aIndex, a := range nums {
		// 取出重复
		if aIndex > 0 && nums[aIndex] == nums[aIndex-1] {
			continue
		}
		// 变为数相加等于某个值
		v := k - a
		//枚举 V
		bIndex := aIndex + 1
		cIndex := n - 1
		// 问题缩减为 两数相加
		for ; bIndex < n; bIndex++ {
			if aIndex > bIndex && nums[bIndex] == nums[bIndex-1] {
				continue
			}
			// 当c指向 + b指向 > 大于结果，C指向想左移动
			for bIndex < cIndex && nums[bIndex]+nums[cIndex] > v {
				cIndex--
			}
			// 两个指针相等
			if bIndex == cIndex {
				break
			}
			// 找到目标值
			if nums[bIndex]+nums[cIndex] == v {
				ans = append(ans, []int{nums[aIndex], nums[bIndex], nums[cIndex]})
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(threeSum([]int{1, 2, -3, 4, -5, 4, 15}, 7))
}
