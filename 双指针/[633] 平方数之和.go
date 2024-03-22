package main

import "math"

//给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a² + b² = c 。
//
//
//
// 示例 1：
//
//
//输入：c = 5
//输出：true
//解释：1 * 1 + 2 * 2 = 5
//
//
// 示例 2：
//
//
//输入：c = 3
//输出：false
//
//
//
//
// 提示：
//
//
// 0 <= c <= 2³¹ - 1
//
//
// Related Topics 数学 双指针 二分查找 👍 450 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func judgeSquareSum(c int) bool {
	startNum := 0
	endNum := int(math.Sqrt(float64(c)))
	for startNum <= endNum {
		sum := startNum*startNum + endNum*endNum
		if sum > c {
			endNum--
		} else if sum < c {
			startNum++
		} else {
			return true
		}
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
