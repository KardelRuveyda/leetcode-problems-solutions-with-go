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
Certainly, here's the explanation with placeholders:

Creating a List (nums):

We created a list called nums, representing stock prices. In this case, we have prices [7, 1, 5, 3, 6, 4].
Calling maxProfit Function:

We called the maxProfit function, which iterates through stock prices to find the best buying and selling points.
Beginning of the Loop and First Step (7):

Starting the loop, we set the initial stock price (7) as minPrice, and initialized res (maximum profit) to zero.
Looping Through Prices:

Within the loop, we check each stock price. If the current price is lower than the current minPrice, we update minPrice as it represents a better buying opportunity.

If the difference between the current price and minPrice is greater than the previously calculated maximum profit (res), we update res.

Result:

After completing the loop, the maxProfit function returns the calculated maximum profit.
main Function:

In the main function, we create an example list of stock prices (nums).
This list is passed to the maxProfit function, and the calculated maximum profit is assigned to the result variable.
Finally, we print the result.
In this specific example, we calculate the maximum profit achievable by determining the best buying and selling points for the stock with prices [7, 1, 5, 3, 6, 4], and print the result.

# Code
```
func maxProfit(prices []int) int {
	minPrice := prices[0]
	res := 0

	for i := range prices {
		if minPrice > prices[i]{
			minPrice = prices[i]
		}else if prices[i] - minPrice >res {
			res = prices[i] - minPrice
		}
	}

	return res;
}
```