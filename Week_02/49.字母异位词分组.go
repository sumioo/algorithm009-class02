/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 *
 * https://leetcode-cn.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (61.47%)
 * Likes:    346
 * Dislikes: 0
 * Total Accepted:    75.7K
 * Total Submissions: 122.8K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
 *
 * 示例:
 *
 * 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
 * 输出:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 *
 * 说明：
 *
 *
 * 所有输入均为小写字母。
 * 不考虑答案输出的顺序。
 *
 *
 */

// @lc code=start
import "sort"
import "strings"

func groupAnagramsA(strs []string) [][]string {
	groups := map[string][]string{}
	for _, str := range strs {
		splits := strings.Split(str, "")
		sort.Strings(splits)
		key := strings.Join(splits, "")
		groups[key] = append(groups[key], str)
	}
	anagrams := [][]string{}
	for _, g := range groups {
		anagrams = append(anagrams, g)
	}
	return anagrams
}

func groupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{}
	for _, str := range strs {
		key := [26]int{}
		for _, c := range str {
			key[c-'a']++
		}
		groups[key] = append(groups[key], str)
	}

	anagrams := [][]string{}
	for _, g := range groups {
		anagrams = append(anagrams, g)
	}
	return anagrams
}

// @lc code=end

