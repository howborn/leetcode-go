package main

//给你一个字符串 s 和一个字符串数组 dictionary ，找出并返回 dictionary 中最长的字符串，该字符串可以通过删除 s 中的某些字符得到。
//
//
// 如果答案不止一个，返回长度最长且字母序最小的字符串。如果答案不存在，则返回空字符串。
//
//
//
// 示例 1：
//
//
//输入：s = "abpcplea", dictionary = ["ale","apple","monkey","plea"]
//输出："apple"
//
//
// 示例 2：
//
//
//输入：s = "abpcplea", dictionary = ["a","b","c"]
//输出："a"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// 1 <= dictionary.length <= 1000
// 1 <= dictionary[i].length <= 1000
// s 和 dictionary[i] 仅由小写英文字母组成
//
//
// Related Topics 数组 双指针 字符串 排序 👍 369 👎 0

func findLongestWord(s string, dictionary []string) string {
	var longestWord string
	var l1, l2 int
	for _, dict := range dictionary {
		// 是否是子串
		if !isWord(s, dict) {
			continue
		}

		l1 = len(longestWord)
		l2 = len(dict)
		// 保留长度最长且字母序最小的子串
		if l1 < l2 || (l1 == l2 && dict < longestWord) {
			longestWord = dict
		}
	}
	return longestWord
}

// 是否是子串
func isWord(s, sub string) bool {
	// 双指针从后遍历
	i, j := len(s)-1, len(sub)-1
	for i >= 0 && j >= 0 {
		// 如果子串的字母相同则移动一位
		if s[i] == sub[j] {
			j--
		}
		i--
	}

	// 子串所有字母遍历完了
	return j == -1
}

//leetcode submit region end(Prohibit modification and deletion)
