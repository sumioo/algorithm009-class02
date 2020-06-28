package week6

import (
	"fmt"
	"sort"
	"strconv"
)

func cutMethodBacktraceCutMethod(n int, collection []int, result *[][]int) {
	if n < 0 {
		return
	}

	if n == 0 {
		*result = append(*result, append([]int{}, collection...))
		return
	}

	// fmt.Println(n)

	for i := 1; i <= n; i++ {
		collection = append(collection, i)
		cutMethodBacktraceCutMethod(n-i, collection, result)
		collection = collection[0 : len(collection)-1]
	}
}

func cutMethod(n int) [][]int {
	result := [][]int{}
	cutMethodBacktraceCutMethod(n, []int{}, &result)
	return result
}

func cutMethodBacktraceCutMethodB(n int, memo map[int][][]int) [][]int {
	fmt.Println(n)

	if n < 0 {
		return nil
	}

	if n == 0 {
		return [][]int{{0}}
	}

	if v, ok := memo[n]; ok {
		return v
	}

	r := [][]int{}
	for i := 1; i <= n; i++ {
		sr := cutMethodBacktraceCutMethodB(n-i, memo)
		for j := range sr {
			x := make([]int, len(sr[j])+1)
			copy(x, sr[j])
			x[len(x)-1] = i
			r = append(r, x)
		}
	}
	memo[n] = r
	return r
}

func cutMethodB(n int) [][]int {
	memo := map[int][][]int{}
	result := cutMethodBacktraceCutMethodB(n, memo)
	// fmt.Println(memo)
	return result
}

func cutMethodBacktraceCutMethodC(n int, collection []int, memo map[int]int, set map[string]bool) int {
	// fmt.Println(n)

	if n < 0 {
		return 0
	}

	if n == 0 {
		fmt.Println(collection)
		x := append([]int{}, collection...)
		sort.Ints(x)
		s := ""
		for i := range x {
			s += strconv.Itoa(x[i])
		}
		if set[s] {
			return 0
		}
		set[s] = true
		return 1
	}

	if v, ok := memo[n]; ok {
		return v
	}

	r := 0
	for i := 1; i <= n; i++ {
		collection = append(collection, i)
		n := cutMethodBacktraceCutMethodC(n-i, collection, memo, set)
		r += n
		collection = collection[:len(collection)-1]

	}
	// memo[n] = r
	return r
}

func cutMethodC(n int) int {
	memo := map[int]int{}
	set := map[string]bool{}
	result := cutMethodBacktraceCutMethodC(n, []int{}, memo, set)
	fmt.Println(memo, set)
	return result
}

func cutRod(plans []int, n int) []int {
	r := make([]int, len(plans))
	r[0] = 0
	for i := 1; i < len(plans); i++ {
		r[i] = plans[i]
		for j := 1; j < i; j++ {
			p := plans[j] + r[i-j] //plans[j]可以替换为r[j].因为他们是等价的，想想它们都可以递归的转化为求解子问题 plans[]， 想象一下当你切下长度为
			if p > r[i] {          //2时，收益为plans[2] 另一段r[i-2]。假如2切割为1，1可获得最大收益，那么在上一次选择切下长度为1时，这种最优切割方案已经
				r[i] = p //被考察了
			}
		}
	}
	return r
}

/*
rn = max(pn,r1 + rn-1, r2 + rn-2,.....,rn-1 + r1)

or

rn = max(pi + rn-i)
1<= i <= n

*/
