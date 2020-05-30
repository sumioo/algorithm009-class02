/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode-cn.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (59.88%)
 * Likes:    195
 * Dislikes: 0
 * Total Accepted:    102.7K
 * Total Submissions: 171.1K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 *
 * 示例 1:
 *
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: s = "rat", t = "car"
 * 输出: false
 *
 * 说明:
 * 你可以假设字符串只包含小写字母。
 *
 * 进阶:
 * 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 *
 */

// @lc code=start
import "sort"
import "strings"

func isAnagramA(s string, t string) bool {
	sSplit := strings.Split(s, "")
	sort.Strings(sSplit)
	s = strings.Join(sSplit, "")
	tSplit := strings.Split(t, "")
	sort.Strings(tSplit)
	t = strings.Join(tSplit, "")
	return s == t
}

/*
test case
"" ""
"aaaa" "aa"
"abc" "abd"
"abc" "bca"
*/
func isAnagramB(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letterCounts := map[byte]int{}
	for i := 0; i < len(s); i++ {
		letterCounts[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if count, ok := letterCounts[t[i]]; ok && count > 0 {
			letterCounts[t[i]]--
		} else {
			return false
		}
	}

	return true

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	letterCounts := [26]int{}

	for i := 0; i < len(s); i++ {
		letterCounts[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		index := t[i] - 'a'
		if letterCounts[index] > 0 {
			letterCounts[index]--
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

