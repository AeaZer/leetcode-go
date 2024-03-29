package olddate

// 100097. 合法分组的最少组数
// https://leetcode.cn/contest/weekly-contest-368/problems/minimum-number-of-groups-to-create-a-valid-assignment/

/*给你一个长度为 n 下标从 0 开始的整数数组 nums 。

我们想将下标进行分组，使得 [0, n - 1] 内所有下标 i 都 恰好 被分到其中一组。

如果以下条件成立，我们说这个分组方案是合法的：

对于每个组 g ，同一组内所有下标在 nums 中对应的数值都相等。
对于任意两个组 g1 和 g2 ，两个组中 下标数量 的 差值不超过 1 。
请你返回一个整数，表示得到一个合法分组方案的 最少 组数。*/

/*func minGroupsForValidAssignment(nums []int) int {
	numsMap := make(map[int]int, 0)
	for i := range nums {
		numsMap[nums[i]]++
	}
	numsCounts := make([]int, 0, len(numsMap))
	smallestCount := math.MaxInt
	for _, value := range numsMap {
		numsCounts = append(numsCounts, value)
		if smallestCount > value {
			smallestCount = value
		}
	}

	var groupSuccess bool
	splitGroupsCount := 1
	var rest bool
	for smallestCount != 0 && !groupSuccess {
		if !rest {
			// 如果数量最少的数恰好可以平分，那么其他元素要么
			for _, count := range numsCounts {
				r := count % (smallestCount + 1)
				if r != smallestCount {
					break
				}
				r = count % (smallestCount)
				if r != 0 && r > count/smallestCount {
					break
				}
			}
			groupSuccess = true
		} else {
			// 如果数量不能平分，那么其他元素
			for _, count := range numsCounts {
				r := count % (smallestCount)
				if r != 0 && r > count/smallestCount {
					break
				}
			}
			groupSuccess = true
		}
		if !groupSuccess {

		}

	}

	var ans int
	for _, count := range numsCounts {
		gs := count / (smallestCount + 1)
		if count%smallestCount == 0 {
			ans += gs
			continue
		}
		ans += gs + 1
	}

	return ans
}*/
