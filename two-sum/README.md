# Intuition
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

 

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
 

Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
 


# Solution 

This is a program written in the Go language, consisting of two main functions: twoSum and main. Let me explain each function in English for someone who is not familiar with the code:

**twoSum Function**
This function takes an array of integers (nums) and a target integer (target). Its purpose is to find a pair of distinct elements in this array whose sum equals the target integer. If found, it returns an integer array containing the indices of these two elements in the array. If no such pair is found, it returns the array [-1, -1].

**The function operates as follows**

- A map named hashMap is created. This map is used to store previously encountered elements and their indices.
- It iterates over all elements in the array (for i, val := range nums).
- For each element, it calculates the result of subtracting that element from the target integer (comp := target - val) and checks if this result has been added to the map before.
- If it has been added (isFound is true), it returns an integer array containing the index of the current element and the index of the other element that adds up to the target integer ([]int{i, idx}).
- If it hasn't been added, the value and index information of the current element are added to the map (hashMap[val] = i).
- If no matching pair is found when the loop completes, it returns the array [-1, -1].

**main Function:**
This function runs an example scenario using the twoSum function. A sample array of integers (nums) and a target integer (target) are specified. The twoSum function is called with these inputs, and the result is printed to the console.

**Overall Program Flow:**
- The main function calls the twoSum function with the nums array and the target target integer.
- The twoSum function checks for a pair of elements in the specified array whose sum equals the target integer.
- The result, if found, is printed to the console by the main function.
- In this example, since the sum of the elements 2 and 7 in the nums array is equal to 9, the twoSum function will return the array [0, 1]. This result will be printed to the console by the main function.

# Code
```
func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)

    for i,val := range nums{
        comp := target-val

        if idx,isFound := hashMap[comp]; isFound{
            return []int{i,idx};
        }

        hashMap[val] = i;
    }
    return []int{-1,-1}
}
```