package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	price := []int{8,7,4,2,8,1,7,7,10,1}
	res := finalPrices(price)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var node *ListNode
	if l1.Val >= l2.Val {
		node = l2
		node.Next = mergeTwoLists(l1, l2.Next)
	} else {
		node = l1
		node.Next = mergeTwoLists(l1.Next, l2)
	}
	return node
}

//26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[count] {
			count++
		}
	}
	return count + 1
}

//27. 删除数组元素
func removeElement(nums []int, val int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

//28. 实现 strStr()
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	if len(haystack) < len(needle) {
		return -1
	}
	var count int
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[count] {
			if count == len(needle) {
				return i - len(needle)
			}
			count++
		} else {
			count = 0
		}
	}
	return -1
}

//两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var isFirst = true
	var head = &ListNode{}
	var tail = &ListNode{}
	var remain = 0
	for l1 != nil || l2 != nil || remain > 0 {
		if l1 != nil {
			remain += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			remain += l2.Val
			l2 = l2.Next
		}
		if isFirst {
			head.Val = remain % 10
			tail = head
			isFirst = false
		} else {
			var node = ListNode{Val: remain % 10}
			tail.Next = &node
			tail = &node
		}
		remain /= 10
	}
	return head
}

//暴力解法
func searchInsert(nums []int, target int) int {
	for index, number := range nums {
		if number >= target {
			return index
		}
	}
	return len(nums)
}

//二分法
func searchInsert2(nums []int, target int) int {
	var left int
	n := len(nums)
	right := n - 1
	for left <= right {
		count := int(math.Ceil(float64(right-left) / 2))
		fmt.Println(count)
		mid := left + count
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}

//53. 最大子序和
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] = nums[i] + nums[i-1]
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}
	return maxSum
}

func lengthOfLastWord(s string) int {
	ss := strings.Split(" ", s)
	fmt.Println(ss)
	return 1

}
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var res map[int]int
	res = make(map[int]int)
	for i, v := range nums2 {
		res[v] = i
	}
	var response []int
	for _, v := range nums1 {
		temp := -1
		for i := res[v]; i < len(nums2); i++ {
			if nums2[i] > v {
				temp = nums2[i]
				break
			}
		}
		response = append(response, temp)
	}
	return response
}

func findJudge(n int, trust [][]int) int {
	exist := map[int]int{}
	for _, numl := range trust {
		exist[numl[0]] = numl[1]
	}
	return 1
}

func plusOne(digits []int) []int {
	lenDigits := len(digits)
	plus := true
	for i := lenDigits - 1; i >= 0; i-- {
		if plus {
			if i == 0 && digits[i] == 9 {
				digits[i] = 0
				digits = append([]int{1}, digits...)
				return digits
			}
			if digits[i] == 9 {
				digits[i] = 0
				plus = true
				continue
			} else {
				digits[i] += 1
				plus = false
			}
		}
	}
	return digits
}

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	//maxLen := 0
	minLen := 0
	if aLen > bLen {
		//maxLen = aLen
		minLen = bLen
	} else {
		//maxLen = bLen
		minLen = aLen
	}
	plus := false
	var res string
	for i := 0; i < minLen; i++ {
		aNum, _ := strconv.Atoi(string(a[i]))
		bNum, _ := strconv.Atoi(string(b[i]))
		if aNum+bNum == 2 {
			if plus {
				res = res + "1"
			}
			plus = true
		} else if aNum+bNum == 1 {
			if plus {
				res = res + "0"
			}
			plus = true
		} else {
			if plus {
				res = res + "1"
			}
			plus = false
		}
	}
	return res
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Len := len(nums1)
	nums2Len := len(nums2)
	var res float64
	if nums1Len == 0 && nums2Len == 0 {
		return res
	}
	var num []int
	if nums1Len == 0 {
		num = nums2
	} else if nums2Len == 0 {
		num = nums1
	} else {
		num = append(nums1, nums2...)
	}
	numLen := len(num)
	if numLen == 1 {
		return float64(num[0])
	}
	//排序
	for i := 0; i < numLen; i++ {
		for j := i + 1; j < numLen; j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}
	if num[0] == num[numLen-1] {
		return float64(num[0])
	}
	if numLen%2 == 1 {
		return float64(num[(numLen-1)/2])
	} else if numLen%2 == 0 {
		index := numLen / 2
		fmt.Print(index)
		return (float64(num[index-1]) + float64(num[index])) / 2
	}
	return res
}

//4,3,2,1,4
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	area := 0
	for {
		var border int
		var weight int
		weight = j - i
		if height[i] > height[j] {
			border = height[j]
		} else {
			border = height[i]
		}
		areaTemp := border * weight
		if area < areaTemp {
			area = areaTemp
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
		if i > j {
			break
		}
	}
	return area
}

//最小因式分解
func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}
	for n >= 4 {
		m := float64(n) / 4
		if m == math.Trunc(m) {
			n = int(m)
		} else {
			return false
		}
	}
	if n != 1 {
		return false
	}
	return true
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	strlen := len(answerKey)
	i := 0
	j := i + 1
	max := 0
	n := 1
	m := 0
	for j < strlen {
		if answerKey[i] == answerKey[j] {
			n++
		} else {
			for m <= k {
				if answerKey[i] == answerKey[j+m] {
					n++
				}
			}
			if answerKey[i] == answerKey[j] {

			}
			if n > max {
				max = n
			}
			n = 1
		}
		i++
		j++
	}
	sum := max + k
	if sum > strlen {
		return strlen
	}
	return sum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}



func preorderTraversal(root *TreeNode) []int {
	var nums []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			nums = append(nums, node.Val)
			dfs(node.Left)
			dfs(node.Right)
		} else {
			return
		}
	}
	dfs(root)
	return nums
}



func inorderTraversal(root *TreeNode) []int {
	var nums []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			nums = append(nums, node.Val)
			return
		} else {
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return nums
}

func finalPrices(prices []int) []int {
	i := 0
	j := i + 1
	var maxPrice []int
	for i < len(prices) {
		if i >= j {
			j++
		}
		if j > len(prices)  - 1 {
			maxPrice = append(maxPrice,  prices[i])
			i++
			j = i + 1
			continue
		}
		if prices[i]  < prices[j] {
			j++
			continue
		} else {
			maxPrice = append(maxPrice,  prices[i] - prices[j])
			i++
			j = i + 1
		}
	}
	return maxPrice
}




/**
88. 合并两个有序数组
给你两个按 非递减顺序 排列的整数数组nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。
示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
示例 2：

输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
解释：需要合并 [1] 和 [] 。
合并结果是 [1] 。
示例 3：

输入：nums1 = [0], m = 0, nums2 = [1], n = 1
输出：[1]
解释：需要合并的数组是 [] 和 [1] 。
合并结果是 [1] 。
注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。
*/
//思路 ： 两数组合起来排序一下，傻子做法
func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[0:m]
	nums2 = nums2[0:n]
	nums1 = append(nums1, nums2...)
	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			if nums1[i] > nums1[j] {
				tmp := nums1[i]
				nums1[i] = nums1[j]
				nums1[j] = tmp
			}
		}
	}
}

/**
169. 多数元素
给定一个大小为 n 的数组nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于⌊  n/2 ⌋的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。
示例1：

输入：nums = [3,2,3]
输出：3
示例2：

输入：nums = [2,2,1,1,1,2,2]
输出：2


提示：
n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109


进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。。
*/
// 傻逼法，先遍历次数存Map，在取出最多的
func majorityElement(nums []int) int {
	n := len(nums)
	var cMap = make(map[int]int)
	for i := 0; i < n; i++ {
		if _, ok := cMap[nums[i]]; ok {
			cMap[nums[i]]++
		} else {
			cMap[nums[i]] = 1
		}
	}
	var maxNum int
	var maxValue int
	for k, v := range cMap {
		if v > maxValue {
			maxValue = v
			maxNum = k
		}
	}
	if maxValue >= n/2 {
		return maxNum
	}
	return 0
}

/*
217. 存在重复元素
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。


示例 1：

输入：nums = [1,2,3,1]
输出：true
示例 2：

输入：nums = [1,2,3,4]
输出：false
示例 3：

输入：nums = [1,1,1,3,3,4,3,2,4,2]
输出：true


提示：

1 <= nums.length <= 105
-109 <= nums[i] <= 109

*/

func containsDuplicate(nums []int) bool {
	n := len(nums)
	var cMap = make(map[int]int)
	for i := 0; i < n; i++ {
		if _, ok := cMap[nums[i]]; ok {
			return true
		} else {
			cMap[nums[i]] = 1
		}
	}
	return false
}

/*
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
*/
// 分别排序然后equals
func isAnagram(s string, t string) bool {
	split := strings.Split(s, "")
	sort.Strings(split)
	s = strings.Join(split, "")

	split1 := strings.Split(t, "")
	sort.Strings(split1)
	t = strings.Join(split1, "")

	return strings.EqualFold(s, t)
}

/*
268. 丢失的数字
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。



示例 1：

输入：nums = [3,0,1]
输出：2
解释：n = 3，因为有 3 个数字，所以所有的数字都在范围 [0,3] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 2：

输入：nums = [0,1]
输出：2
解释：n = 2，因为有 2 个数字，所以所有的数字都在范围 [0,2] 内。2 是丢失的数字，因为它没有出现在 nums 中。
示例 3：

输入：nums = [9,6,4,2,3,5,7,0,1]
输出：8
解释：n = 9，因为有 9 个数字，所以所有的数字都在范围 [0,9] 内。8 是丢失的数字，因为它没有出现在 nums 中。
示例 4：

输入：nums = [0]
输出：1
解释：n = 1，因为有 1 个数字，所以所有的数字都在范围 [0,1] 内。1 是丢失的数字，因为它没有出现在 nums 中。


提示：

n == nums.length
1 <= n <= 104
0 <= nums[i] <= n
nums 中的所有数字都 独一无二


进阶：你能否实现线性时间复杂度、仅使用额外常数空间的算法解决此问题?
*/

func missingNumber(nums []int) int {
	//nums := []int{0, 1, 2, 4, 5}
	lens := len(nums)
	for i := 0; i < lens; i++ {
		for j := i + 1; j < lens; j++ {
			if nums[i] > nums[j] {
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
		}
	}
	last := lens - 1
	if nums[last] != lens {
		return lens
	}
	if nums[0] != 0 {
		return 0
	}
	k := 0
	for s := 1; s < lens+1; s++ {
		if nums[s]-nums[s-1] > 1 {
			k = nums[s-1] + 1
			break
		}
	}
	return k
}

/*
349. 两个数组的交集
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]
示例 2：

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的


提示：

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000

*/

func intersection(nums1 []int, nums2 []int) []int {
	hash1 := make(map[int]int)
	hash2 := make(map[int]int)

	for _, i := range nums1 {
		hash1[i] = 1
	}
	for _, j := range nums2 {
		hash2[j] = 1
	}
	var res []int
	for d, _ := range hash1 {
		_, b := hash2[d]
		if b != false {
			res = append(res, d)
		}
	}
	return res
}

/**
350. 两个数组的交集 II
给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。



示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]


提示：

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000


进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小，哪种方法更优？
如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/

func intersect(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)

	for _, i := range nums2 {
		_, b := hash[i]
		if b == false {
			hash[i] = 1
		} else {
			hash[i]++
		}
	}
	var res []int
	for _, nums := range nums1 {
		v, c := hash[nums]
		if c != false && v > 0 {
			res = append(res, nums)
			hash[nums]--
		}
	}
	return res
}

/**
389. 找不同
给定两个字符串 s 和 t ，它们只包含小写字母。

字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。

请找出在 t 中被添加的字母。



示例 1：

输入：s = "abcd", t = "abcde"
输出："e"
解释：'e' 是那个被添加的字母。
示例 2：

输入：s = "", t = "y"
输出："y"


提示：

0 <= s.length <= 1000
t.length == s.length + 1
s 和 t 只包含小写字母
*/

func findTheDifference(s string, t string) byte {
	var res int32
	for _, v := range s {
		res -= v
	}
	for _, k := range t {
		res += k
	}
	return byte(res)
}

/**
414. 第三大的数
给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。



示例 1：

输入：[3, 2, 1]
输出：1
解释：第三大的数是 1 。
示例 2：

输入：[1, 2]
输出：2
解释：第三大的数不存在, 所以返回最大的数 2 。
示例 3：

输入：[2, 2, 3, 1]
输出：1
解释：注意，要求返回第三大的数，是指在所有不同数字中排第三大的数。
此例中存在两个值为 2 的数，它们都排第二。在所有不同数字中排第三大的数为 1 。


提示：

1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1


进阶：你能设计一个时间复杂度 O(n) 的解决方案吗？


*/

func thirdMax(nums []int) int {
	map1 := make(map[int]int)
	//去重
	for _, v := range nums {
		map1[v] = 1
	}
	var nums1 []int
	for k, _ := range map1 {
		nums1 = append(nums1, k)
	}
	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			if nums1[i] < nums1[j] {
				tmp := nums1[i]
				nums1[i] = nums1[j]
				nums1[j] = tmp
			}
		}
	}
	if len(nums1) < 3 {
		return nums1[0]
	}
	return nums1[2]
}

/**
455. 分发饼干
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。

对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。


示例 1:

输入: g = [1,2,3], s = [1,1]
输出: 1
解释:
你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
所以你应该输出1。
示例 2:

输入: g = [1,2], s = [1,2,3]
输出: 2
解释:
你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
你拥有的饼干数量和尺寸都足以让所有孩子满足。
所以你应该输出2.


提示：

1 <= g.length <= 3 * 104
0 <= s.length <= 3 * 104
1 <= g[i], s[j] <= 231 - 1

*/

func findContentChildren(g []int, s []int) int {
	for i := 0; i < len(g); i++ {
		for j := i + 1; j < len(g); j++ {
			if g[i] > g[j] {
				tmp := g[i]
				g[i] = g[j]
				g[j] = tmp
			}
		}
	}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				tmp := s[i]
				s[i] = s[j]
				s[j] = tmp
			}
		}
	}

	i := 0
	j := 0
	ch := 0

	for i < len(s) && j < len(g) {
		if s[i] >= g[j] {
			ch++
			i++
			j++
		} else {
			i++
		}
	}
	return ch
}

/**
506. 相对名次
给你一个长度为 n 的整数数组 score ，其中 score[i] 是第 i 位运动员在比赛中的得分。所有得分都 互不相同 。

运动员将根据得分 决定名次 ，其中名次第 1 的运动员得分最高，名次第 2 的运动员得分第 2 高，依此类推。运动员的名次决定了他们的获奖情况：

名次第 1 的运动员获金牌 "Gold Medal" 。
名次第 2 的运动员获银牌 "Silver Medal" 。
名次第 3 的运动员获铜牌 "Bronze Medal" 。
从名次第 4 到第 n 的运动员，只能获得他们的名次编号（即，名次第 x 的运动员获得编号 "x"）。
使用长度为 n 的数组 answer 返回获奖，其中 answer[i] 是第 i 位运动员的获奖情况。



示例 1：

输入：score = [5,4,3,2,1]
输出：["Gold Medal","Silver Medal","Bronze Medal","4","5"]
解释：名次为 [1st, 2nd, 3rd, 4th, 5th] 。
示例 2：

输入：score = [10,3,8,9,4]
输出：["Gold Medal","5","Bronze Medal","Silver Medal","4"]
解释：名次为 [1st, 5th, 3rd, 2nd, 4th] 。
*/
//
//func findRelativeRanks(score []int) []string {
//
//}

/**
1859. 将句子排序
一个 句子 指的是一个序列的单词用单个空格连接起来，且开头和结尾没有任何空格。每个单词都只包含小写或大写英文字母。

我们可以给一个句子添加 从 1 开始的单词位置索引 ，并且将句子中所有单词 打乱顺序 。

比方说，句子 "This is a sentence" 可以被打乱顺序得到 "sentence4 a3 is2 This1" 或者 "is2 sentence4 This1 a3" 。
给你一个 打乱顺序 的句子 s ，它包含的单词不超过 9 个，请你重新构造并得到原本顺序的句子。



示例 1：

输入：s = "is2 sentence4 This1 a3"
输出："This is a sentence"
解释：将 s 中的单词按照初始位置排序，得到 "This1 is2 a3 sentence4" ，然后删除数字。
示例 2：

输入：s = "Myself2 Me1 I4 and3"
输出："Me Myself and I"
解释：将 s 中的单词按照初始位置排序，得到 "Me1 Myself2 and3 I4" ，然后删除数字。


提示：

2 <= s.length <= 200
s 只包含小写和大写英文字母、空格以及从 1 到 9 的数字。
s 中单词数目为 1 到 9 个。
s 中的单词由单个空格分隔。
s 不包含任何前导或者后缀空格。
*/
func sortSentence(s string)  string{
	strArrayNew := strings.Split(s, " ")
	var cMap = make(map[int]string)
	var temArr []int
	for _, s1 := range strArrayNew {
		tmpS := s1[0:len(s1)-1]
		tmpO := s1[len(s1)-1:]
		tmpI, _ := strconv.Atoi(tmpO)
		cMap[tmpI] = tmpS
		temArr = append(temArr, tmpI)
	}
	for i := 0; i < len(temArr); i++ {
		for j := i + 1; j < len(temArr); j++ {
			if temArr[i] > temArr[j] {
				tmp := temArr[i]
				temArr[i] = temArr[j]
				temArr[j] = tmp
			}
		}
	}
	var res []string
	for _, v := range temArr{
		res = append(res, cMap[v])
	}
	response := strings.Join(res, " ")
	return response
}


/**
2126. 摧毁小行星
给你一个整数 mass ，它表示一颗行星的初始质量。再给你一个整数数组 asteroids ，其中 asteroids[i] 是第 i 颗小行星的质量。

你可以按 任意顺序 重新安排小行星的顺序，然后让行星跟它们发生碰撞。如果行星碰撞时的质量 大于等于 小行星的质量，那么小行星被 摧毁 ，并且行星会 获得 这颗小行星的质量。否则，行星将被摧毁。

如果所有小行星 都 能被摧毁，请返回 true ，否则返回 false 。



示例 1：

输入：mass = 10, asteroids = [3,9,19,5,21]
输出：true
解释：一种安排小行星的方式为 [9,19,5,3,21] ：
- 行星与质量为 9 的小行星碰撞。新的行星质量为：10 + 9 = 19
- 行星与质量为 19 的小行星碰撞。新的行星质量为：19 + 19 = 38
- 行星与质量为 5 的小行星碰撞。新的行星质量为：38 + 5 = 43
- 行星与质量为 3 的小行星碰撞。新的行星质量为：43 + 3 = 46
- 行星与质量为 21 的小行星碰撞。新的行星质量为：46 + 21 = 67
所有小行星都被摧毁。
示例 2：

输入：mass = 5, asteroids = [4,9,23,4]
输出：false
解释：
行星无论如何没法获得足够质量去摧毁质量为 23 的小行星。
行星把别的小行星摧毁后，质量为 5 + 4 + 9 + 4 = 22 。
它比 23 小，所以无法摧毁最后一颗小行星。
*/

func asteroidsDestroyed(mass int, asteroids []int) bool {
	//for i := 0; i < len(asteroids); i++ {
	//	for j := i + 1; j < len(asteroids); j++ {
	//		if asteroids[i] > asteroids[j] {
	//			tmp := asteroids[i]
	//			asteroids[i] = asteroids[j]
	//			asteroids[j] = tmp
	//		}
	//	}
	//}
	sort.Ints(asteroids)
	for _, v := range asteroids {
		if mass >= v {
			mass += v
		} else {
			return false
		}
	}
	return true
}


func maxSubArray1(nums []int) int {
	res := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := nums[i] + max
		if tmp > 0 {
			if tmp > res {
				res = tmp
				max = tmp
			} else {
				max = nums[i]
			}
		} else {
			max = nums[i]
		}
	}
	return max
}