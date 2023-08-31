## Problem

You are given an array prices where prices[i] is the 
price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one 
stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction.
If you cannot achieve any profit, return 0.
 

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

## Notes

Initial thoughts:
- Probably want some way of recording the price that we bought it for
and what maximum sell price we would be able to achieve with that buy price
- The maximum sell price will always be the maximum value in the list, but can
we come across that in less than O(n) time? prob not, we already need to walk
it.

What we're dealing with here could be thought of as a tree. Each node is a price
and the leaves are the differences with all the other nodes

So, for list = [7,1,5,3,6,4]

list - 7 = [0, -6, -2, -4, -1, -3]

All of these are negative, so 7 definitely isn't a good buy candidate

Numbers can only be associated with numbers coming ahead of them, so we only
need to look forward.

I'm going to try this solution:
1. Starting at index zero, calculate `num - list` and find the biggest number
2. Record that biggest number in a placeholder variable
3. Continue through the list, updating the placeholder if the biggest number
exceeds the current.
4. Return the placeholder biggest number

Immediately there could be an optimization for large amount of repeating numbers
by using a hashmap, the hashmap could serve as a cache.

E.g. for [7, 1, 5, 1, 3, 6, 4]. The second 1 could use the computation from
the first 1.

Well... the solution above resulted in a *Time limit exceeded* warning XD

Let's see if we can do it without using a nested for loop

What we're looking for is the biggest number - smallest number that came
before it.

What if we have two placeholders?
1. For smallest number
2. For biggest number so far?

Hmm.. 

What about sorting the array first and note the original indices?

[7,1,5,3,6,4] -> [1,3,4,5,6,7]

arr =       [1,3,4,5,6,7]
indices =   [1,3,5,2,4,0]

After its sorted, we have a new problem, find the greatest distance between
increasing indices.

But, it's a similar pattern, I need to apply a set of different conditions
on all the members of the array to see which numbers are compatible with
each other.

### Looked at solution

After looking at solution I was on the right track, to keep a record of the
minimum value.

That solution works, but my performance is still super crap compared to other
users. 

Worked on it a little more, one optimization is to wrap the update of the 
largest difference in an else if statement, since if we're updating the smallest
element, we definitely don't need to update the difference.
