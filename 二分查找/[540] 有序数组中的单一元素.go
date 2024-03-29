package main

//给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
//
// 请你找出并返回只出现一次的那个数。
//
// 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
//
//
//
// 示例 1:
//
//
//输入: nums = [1,1,2,3,3,4,4,8,8]
//输出: 2
//
//
// 示例 2:
//
//
//输入: nums =  [3,3,7,7,10,11,11]
//输出: 10
//
//
//
//
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 10⁵
// 0 <= nums[i] <= 10⁵
//
//
// Related Topics 数组 二分查找 👍 662 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 插入的数据会打断配对的数组序列，正常不打乱的情况下：
// 中位数为奇数则跟前面一个数为一对
// 反之则跟后面的数为一对
// 若这一对数据不相等则在左区间有单独的数插入，反之则在右区间
// [1,1,2,2,3,3,4,8,8]
// [3,3,7,7,10,11,11]
// [3,3,7,7,10]
func singleNonDuplicate(nums []int) int {
	var mid int
	left, right := 0, len(nums)
	// 开区间
	for left < right {
		mid = (left + right) / 2
		// 中间位是奇数时跟前一个为一对
		if mid%2 == 1 {
			// 统一找到一对数的最左边的数
			mid--
		}

		// 如果刚好能匹配一对那单独的数就在右区间
		if nums[mid] == nums[mid+1] {
			left = mid + 2
		} else {
			right = mid - 1
		}
	}

	return nums[left]
}

//leetcode submit region end(Prohibit modification and deletion)
