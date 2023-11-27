package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	hash := make(map[int]int)

	for _, val := range nums {
		if _, isFound := hash[val]; isFound {
			return true
		}
		hash[val] = 1
	}
	return false
}

func main() {
	nums := []int{1, 2, 2, 3}
	result := containsDuplicate(nums)
	fmt.Println(result)
}
