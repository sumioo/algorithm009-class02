package week8

import "fmt"

/*
 *
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (48.62%)
 * Likes:    1013
 * Dislikes: 0
 * Total Accepted:    197.9K
 * Total Submissions: 406K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 *
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 *
 * 注意：给定 n 是一个正整数。
 *
 * 示例 1：
 *
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 *
 * 示例 2：
 *
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 *
 *
 */

//递归 memo

func helper(n int, cache map[int]int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	//query cache
	if v, ok := cache[n]; ok {
		return v
	}

	//cal result
	r := helper(n-1, cache) + helper(n-2, cache)
	//cache
	cache[n] = r
	return r
}

func climbStairsR(n int) int {
	if n <= 0 {
		return 0
	}
	cache := map[int]int{}
	return helper(n, cache)
}

func climbStairsBU(n int) int {
	if n < 0 {
		return 0
	}

	dp := make([]int, n)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[n] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func climbStairsBUoP(n int) int {
	if n < 0 {
		return 0
	}

	if n <= 2 {
		return n
	}

	oneStepDown := 2 //对于第3来说
	twoStepDown := 1
	r := 0

	for i := 3; i <= n; i++ {
		r = oneStepDown + twoStepDown
		twoStepDown = oneStepDown
		oneStepDown = r
	}
	return r
}

func climbeStairsBacktraceLimit(n int, steps []int, prevChoice int, cache map[[2]int]int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	if v, ok := cache[[2]int{n, prevChoice}]; ok {
		return v
	}

	count := 0
	for _, step := range steps {
		if step != prevChoice {
			count += climbeStairsBacktraceLimit(n-step, steps, step, cache)
		}
	}

	cache[[2]int{n, prevChoice}] = count
	return count
}

func climbStairsLimitChoice(n int, steps []int) int {
	cache := map[[2]int]int{}
	result := climbeStairsBacktraceLimit(n, steps, 0, cache)
	return result
}

//假如支持任意steps的话，基准条件很难生成，还不如用递归的办法
func climbStairsLimitChoiceBU(n int) int {
	//steps 1, 2, 3
	//这里状态定义的是上一阶段跨了几步到达该阶段，如果是跨了1步，那么该阶段的这个状态只能从2，3转移得来
	dp := make([][3]int, n+1)
	dp[0] = [3]int{1, 0, 0} //第一个格子表示跨了1步到达，第二个格子表示跨了2步到达，第三个格子表示跨了3步到达，
	dp[1] = [3]int{1, 0, 0}
	dp[2] = [3]int{0, 1, 0}
	dp[3] = [3]int{1, 1, 1}

	for i := 4; i <= n; i++ {
		dp[i][0] = dp[i-1][1] + dp[i-1][2]
		dp[i][1] = dp[i-2][0] + dp[i-2][2]
		dp[i][2] = dp[i-3][0] + dp[i-3][1]
	}

	fmt.Println(dp)
	return dp[n][0] + dp[n][1] + dp[n][2]
}

//打印出所有方案

func climbeStairsPathBacktrace(n int, steps []int, cache map[int][][]int) [][]int {
	if n < 0 {
		return nil
	}
	if n == 0 {
		return [][]int{[]int{}}
	}

	if v, ok := cache[n]; ok {
		return v
	}

	paths := [][]int{}
	for _, step := range steps {
		results := climbeStairsPathBacktrace(n-step, steps, cache)
		if results != nil {
			for _, result := range results {
				path := make([]int, len(result))
				copy(path, result)
				path = append(path, step)
				paths = append(paths, path)

			}
		}

	}
	cache[n] = paths
	return paths
}

func climbStairsPaths(n int, steps []int) [][]int {
	cache := map[int][][]int{}
	results := climbeStairsPathBacktrace(n, steps, cache)
	return results
}

func climbeStairsPathBacktraceLimit(n int, steps []int, prevChoice int, cache map[[2]int][][]int) [][]int {
	if n < 0 {
		return nil
	}
	if n == 0 {
		return [][]int{[]int{}}
	}

	if v, ok := cache[[2]int{n, prevChoice}]; ok {
		fmt.Println("found at cache", n, prevChoice)
		return v
	}

	fmt.Println("new", n, prevChoice)
	paths := [][]int{}
	for _, step := range steps {
		if step != prevChoice {
			results := climbeStairsPathBacktraceLimit(n-step, steps, step, cache)
			for _, result := range results {
				path := make([]int, len(result))
				copy(path, result)
				path = append(path, step)
				paths = append(paths, path)
			}
		}
	}
	cache[[2]int{n, prevChoice}] = paths
	return paths
}

func climbStairsPathsLimitChoice(n int, steps []int) [][]int {
	cache := map[[2]int][][]int{}
	results := climbeStairsPathBacktraceLimit(n, steps, 0, cache)
	return results
}
