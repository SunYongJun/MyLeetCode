package leetcode

/**

4. 寻找两个正序数组的中位数
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。

进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？



示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
示例 4：

输入：nums1 = [], nums2 = [1]
输出：1.00000
示例 5：

输入：nums1 = [2], nums2 = []
输出：2.00000


提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106

*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i,j := 0,0
	l1,l2 := len(nums1),len(nums2)
	nums3 := make([]int,0)

	for i < l1 || j < l2 {

		if j == l2 || i < l1 && nums1[i] < nums2[j] {
			nums3 = append(nums3,nums1[i])
			i++
		}else if i == l1 || j < l2 && nums1[i] >= nums2[j] {
			nums3 = append(nums3,nums2[j])
			j++
		}
	}
	l3 := len(nums3)
	if l3 % 2 == 0 {
		return (float64( nums3[l3 / 2] ) + float64(nums3[l3 / 2 - 1])) / 2
	}else{
		return float64( nums3[l3 / 2] )
	}

}