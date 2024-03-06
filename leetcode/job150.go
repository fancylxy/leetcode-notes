package main

import "fmt"

func main() {
	//n1 := []int{1, 2, 3, 0, 0, 0}
	n2 := []int{1,2,3, 4,5}
	 rotate(n2,2)
}

func merged(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m - 1, n - 1, m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		k--
		j--
	}
	fmt.Println(nums1)
}

//27. 删除数组元素
func removeElement2(nums []int, val int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

func removeDuplicates2(nums []int) int {
	count := 0
	dup := make(map[int]int)
	for i :=0; i < len(nums); i++ {
		if dup[nums[i]] <= 0 {
			nums[count] = nums[i]
			dup[nums[i]]++
			count++
		}
	}
	return count
}

func majorityElement2(nums []int) int {
	ll := len(nums)
	dup := make(map[int]int)
	for i := 0; i < ll; i++ {
		dup[nums[i]]++
	}
	for j,k := range dup {
		if float64(k) >= float64(ll) / float64(2) {
			return j
		}
	}
	return 0
}

func rotate(nums []int, k int)  {
	ll := len(nums)
	numsCopy := append(nums[ll - k:])
	numsCopy2 := append(nums[:ll - k])
	fmt.Println(numsCopy)
	fmt.Println(numsCopy2)

	copy(nums, numsCopy)
	copy(nums, numsCopy2)
	fmt.Println(nums)
}