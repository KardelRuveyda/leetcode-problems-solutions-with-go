package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, val := range nums {
		comp := target - val

		if idx, isFound := hashMap[comp]; isFound {
			return []int{i, idx}
		}

		hashMap[val] = i
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result)
}
