/*
Given a list arr of n integers, return sums of all subsets in it. Output sums can be printed in any order.

Example--

Input:
n = 2
arr[] = {2, 3}
Output:
0 2 3 5
Explanation:
When no elements is taken then Sum = 0.
When only 2 is taken then Sum = 2.
When only 3 is taken then Sum = 3.
When element 2 and 3 are taken then 
Sum = 2+3 = 5.
*/


package main

import (
	"fmt"
	"sort"
)

func subsetSums(arr []int) []int {
	res := []int{}
	var backtrack func(int, int)
	backtrack = func(cur int, sum int) {
		if cur >= len(arr) {
			res = append(res, sum)
			return
		}
		// to include in sum
		backtrack(cur+1, sum+arr[cur])
		// to not include in sum
		backtrack(cur+1, sum)
	}
	backtrack(0, 0)
	sort.Ints(res)
	return res
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(subsetSums(arr))
}

/*
Here's a breakdown of the Go code:

- `subsetSums(arr []int) []int`: This is the main function that takes an array of integers as input and returns a slice of integers. The returned slice contains the sums of all possible subsets of the input array.

- `res := []int{}`: This line initializes `res`, which is a slice that will store the sums of all possible subsets.

- `var backtrack func(int, int)`: This line declares the `backtrack` function, which is a recursive function that generates all possible subsets and their sums.

- `backtrack = func(cur int, sum int) {...}`: This block of code defines the `backtrack` function. It takes two integers as input: `cur` (the current index in `arr`) and `sum` (the sum of the current subset).

- `if cur >= len(arr) {...}`: This `if` statement checks if `cur` has reached the end of `arr`. If it has, it adds `sum` to `res` and returns.

- `backtrack(cur+1, sum+arr[cur])` and `backtrack(cur+1, sum)`: These lines are the recursive calls to `backtrack`. The first call includes the current number in the sum, and the second call does not include the current number in the sum.

- `sort.Ints(res)`: This line sorts `res` in ascending order.

- `func main() {...}`: This is the main function that runs when the program starts. It initializes `arr` and prints the result of `subsetSums(arr)`.


*/
