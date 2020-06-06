/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (53.51%)
 * Likes:    736
 * Dislikes: 0
 * Total Accepted:    117.7K
 * Total Submissions: 219.5K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */

// @lc code=start

/*
重新描述下题目
有两个位置，从两个字符串中各选取一个字符填入这个位置，求可能的组合
*/

func backtrace(index int, digits string, letters map[string]string, collection string, result *[]string) {
	if index == len(digits) {
		*result = append(*result, collection)
		return
	}

	for _, r := range letters[string(digits[index])] {
		backtrace(index+1, digits, letters, collection+string(r), result)
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	letters := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	result := []string{}
	backtrace(0, digits, letters, "", &result)
	return result
}

// @lc code=end

