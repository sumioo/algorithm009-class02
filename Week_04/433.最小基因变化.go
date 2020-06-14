/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 *
 * https://leetcode-cn.com/problems/minimum-genetic-mutation/description/
 *
 * algorithms
 * Medium (50.38%)
 * Likes:    37
 * Dislikes: 0
 * Total Accepted:    5K
 * Total Submissions: 9.9K
 * Testcase Example:  '"AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]'
 *
 * 一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。
 *
 * 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。
 *
 * 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。
 *
 * 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。
 *
 * 现在给定3个参数 — start, end,
 * bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回
 * -1。
 *
 * 注意:
 *
 *
 * 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
 * 所有的目标基因序列必须是合法的。
 * 假定起始基因序列与目标基因序列是不一样的。
 *
 *
 * 示例 1:
 *
 *
 * start: "AACCGGTT"
 * end:   "AACCGGTA"
 * bank: ["AACCGGTA"]
 *
 * 返回值: 1
 *
 *
 * 示例 2:
 *
 *
 * start: "AACCGGTT"
 * end:   "AAACGGTA"
 * bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
 *
 * 返回值: 2
 *
 *
 * 示例 3:
 *
 *
 * start: "AAAAACCC"
 * end:   "AACCCCCC"
 * bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
 *
 * 返回值: 3
 *
 *
 */

// @lc code=start

/*
解题的关键是如何获取顶点的边，根据题意可以得知如果基因改变任意一个基因的后并且落在基因库的话，
则该改变后的基因为一条边，所以我们可以枚举出所有组合看看这些组合是否落在基因库，但这样的话将会话费很多时间，
那么我们可以从基因库出发，看看基因库的基因是否和顶点只有一个基因的差别。
*/

func isOneGeneDiff(geneA string, geneB string) bool {
	n := len(geneA)
	for i := 0; i < n; i++ {
		if geneA[i] != geneB[i] {
			return geneA[i+1:] == geneB[i+1:]
		}
	}
	return false
}

func searchEdges(gene string, bank []string) []string {
	edges := []string{}
	for _, g := range bank {
		if isOneGeneDiff(gene, g) {
			edges = append(edges, g)
		}
	}
	return edges
}

func minMutation(start string, end string, bank []string) int {
	level := 0
	queue := []string{}
	visited := map[string]bool{}
	queue = append(queue, start)
	visited[start] = true
	found := false

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			vertex := queue[0]
			queue = queue[1:]
			if vertex == end {
				found = true
				goto RETURN
			}

			for _, edge := range searchEdges(vertex, bank) {
				if !visited[edge] {
					queue = append(queue, edge)
					visited[edge] = true
				}
			}
		}
		level++
	}

RETURN:
	if found {
		return level
	}
	return -1
}

// @lc code=end

