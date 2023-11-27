# Question
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

 

Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
 

Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
# Solution
This Go program is designed to check whether there are duplicate elements in an array. The program utilizes a hash map to efficiently detect repeated elements.

Here's a step-by-step explanation of the code:

containsDuplicate Function:

The containsDuplicate function takes an array of numbers (nums) and checks whether there are duplicate elements in the array.
A hash map named hash is created to keep track of whether numbers in the array are repeated.
A loop iterates through each element in the array:
If the element is already present in the hash map (i.e., it has been seen before), the function returns true, indicating that a duplicate element is found.
If the element is not seen before, it is added to the hash map with a value indicating its presence.
main Function:

The main function creates an example array of numbers (nums).
This array is passed to the containsDuplicate function, and the result returned by the function is printed to the console.
Example:

For instance, the nums array is given as [1, 2, 2, 3]. Since there is a repeated element (2) in this array, the program returns true.
In summary, the program checks for duplicate elements in an array using a hash map and prints the result. This type of check is useful, especially for verifying the uniqueness of elements in an array.

# Code
```
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
```