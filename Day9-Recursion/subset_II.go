/*
SubsetII

Given an integer array nums that may contain duplicates, return all possible 
subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:

Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
Example 2:

Input: nums = [0]
Output: [[],[0]]
*/

package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	subset := []int{}

	var backtrack func(int)
	backtrack = func(cur int) {
		if cur >= len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}

		// to include
		subset = append(subset, nums[cur])
		backtrack(cur + 1)

		// not include
		subset = subset[:len(subset)-1]

		// skip duplicates
		for cur+1 < len(nums) && nums[cur] == nums[cur+1] {
			cur++
		}

		backtrack(cur + 1)
	}

	backtrack(0)

	return res
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

/*
Here's a breakdown of the Go code:

- `subsetsWithDup(nums []int) [][]int`: This is the main function that takes a slice of integers as input and returns a slice of slices of integers. The returned slice contains all possible subsets of the input slice.

- `sort.Ints(nums)`: This line sorts `nums` to ensure that duplicate elements are adjacent.

- `res := [][]int{}`: This line initializes `res`, which is a slice that will store all subsets.

- `subset := []int{}`: This line initializes `subset`, which is a slice that will store the current subset.

- `var backtrack func(int)`: This line declares the `backtrack` function, which is a recursive function that generates all possible subsets.

- `backtrack = func(cur int) {...}`: This block of code defines the `backtrack` function. It takes one integer as input: `cur` (the current index in `nums`).

- `if cur >= len(nums) {...}`: This `if` statement checks if `cur` has reached the end of `nums`. If it has, it adds `subset` to `res` and returns.

- `subset = append(subset, nums[cur])` and `subset = subset[:len(subset)-1]`: These lines are the recursive calls to `backtrack`. The first call includes the current number in the subset, and the second call does not include the current number in the subset.

- `for cur+1 < len(nums) && nums[cur] == nums[cur+1] {...}`: This `for` loop skips all duplicate elements.

- `func main() {...}`: This is the main function that runs when the program starts. It initializes `nums` and prints the result of `subsetsWithDup(nums)`.

*/