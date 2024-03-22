package main

//给你一个字符串 s，最多 可以从中删除一个字符。
//
// 请你判断 s 是否能成为回文字符串：如果能，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
//输入：s = "aba"
//输出：true
//
//
// 示例 2：
//
//
//输入：s = "abca"
//输出：true
//解释：你可以删除字符 'c' 。
//
//
// 示例 3：
//
//
//输入：s = "abc"
//输出：false
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁵
// s 由小写英文字母组成
//
//
// Related Topics 贪心 双指针 字符串 👍 633 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func validPalindrome(s string) bool {
	startIndex := 0
	endIndex := len(s) - 1
	for ; startIndex < endIndex; startIndex, endIndex = startIndex+1, endIndex-1 {
		if s[startIndex] != s[endIndex] {
			return isPalindromeSub(s, startIndex+1, endIndex) || isPalindromeSub(s, startIndex, endIndex-1)
		}
	}

	return true
}

func isPalindromeSub(s string, i, j int) bool {
	for ; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
