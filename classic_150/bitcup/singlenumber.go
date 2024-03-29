package bitcup

// 136. 只出现一次的数字
// https://leetcode.cn/problems/single-number/description/?envType=study-plan-v2&envId=top-interview-150
// 给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
// 你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

func singleNumber(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	for num := range m {
		if m[num] == 1 {
			return num
		}
	}
	return -1
}

// singleNumber2 异或运算
// 任何数和 000 做异或运算
// 任何数和其自身做异或运算
// 异或运算满足交换律和结合律
func singleNumber2(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
