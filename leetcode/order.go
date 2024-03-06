package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	price := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	fmt.Println(threeSumClosest(price, -2))
}

func findRelativeRanks(score []int) []string {
	getOrderName := func(idx int) string {
		awardMap := map[int]string{
			0: "Gold Medal",
			1: "Silver Medal",
			2: "Bronze Medal",
		}
		if idx <= 2 {
			return awardMap[idx]
		}
		return strconv.Itoa(idx + 1)
	}
	//原始排序
	ori := make([]int, len(score))
	copy(ori, score)

	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score); j++ {
			if score[i] > score[j] {
				k := score[i]
				score[i] = score[j]
				score[j] = k
			}
		}
	}
	orderName := make(map[int]string)
	for index, value := range score {
		orderName[value] = getOrderName(index)
	}
	var order []string
	for _, v := range ori {
		order = append(order, orderName[v])
	}
	return order
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for k, v := range nums {
		if k%2 != 0 {
			sum += v
		}
	}
	return sum
}

func findLHS(nums []int) int {
	sort.Ints(nums)
	//去重
	var oneArr []int
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v]++
	}
	for k, _ := range countMap {
		oneArr = append(oneArr, k)
	}
	sort.Ints(oneArr)

	sum := 0
	for i := 0; i < len(oneArr)-1; i++ {
		if oneArr[i+1]-oneArr[i] <= 1 {
			tmp := countMap[oneArr[i]] + countMap[oneArr[i+1]]
			if tmp > sum {
				fmt.Println(i)
				fmt.Println(tmp)
				sum = tmp
			}
		}
	}
	return sum
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	lens := len(nums)
	oneMap := make(map[string]int)
	var res [][]int
	for i := 0; i < lens; i++ {
		end := lens - 1
		j := i + 1
		for end > j {
			sum := nums[i] + nums[j] + nums[end]
			if sum > 0 {
				end--
			}
			if sum < 0 {
				j++
			}
			if sum == 0 {
				tmp := []int{nums[i], nums[j], nums[end]}
				oneKey := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j]) + strconv.Itoa(nums[end])
				_, ok := oneMap[oneKey]
				if !ok {
					res = append(res, tmp)
					oneMap[oneKey] = 1
				}
				end--
				j++
			}
		}
	}
	return res
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	n := len(nums)
	sum := 0
	min := math.MaxFloat64
	for i := 0; i < n; i++ {
		end := n - 1
		start := i + 1
		for end > start {

			sumTmp := nums[i] + nums[start] + nums[end]
			dxTmp := math.Abs(float64(target - sumTmp))

			if sumTmp == target {
				return target
			}

			if dxTmp <= min {
				sum = sumTmp
				min = dxTmp
			}

			if sum > target {
				end--
			} else {
				start++
			}
		}
	}
	return sum
}
