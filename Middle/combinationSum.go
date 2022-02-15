package main

func combinationSum(candidates []int, target int) (ans [][]int) {
	var comb []int
	length := len(candidates)
	var dfs func(target, index int)
	dfs = func(target, index int) {
		// 这个地方剪枝，因为每一个元素都是大于 0 注：题目要求：1 <= candidates[i] <= 200
		if index == length || target < 0 {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}

		for i, element := range candidates {
			// 不可以倒过去寻值，即可以[2,2,3]，不可以[2,3,2]
			if index > i {
				continue
			}
			comb = append(comb, element)
			dfs(target-element, i)
			if len(comb) > 1 {
				comb = comb[:len(comb)-1]
			} else {
				comb = []int{}
			}
		}
	}
	dfs(target, 0)
	return
}
