# Question

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

 

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
 

Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 104

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
