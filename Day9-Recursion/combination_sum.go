/*

Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the 
frequency
 of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

 

Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
*/

package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    backtrack(candidates, target, &res, []int{}, 0, 0)
    return res
}

func backtrack(candidates []int, target int, res *[][]int, temp []int, cur int, total int) {
    if total == target {
        combination := make([]int, len(temp))
        copy(combination, temp)
        *res = append(*res, combination)
        return
    }
    if cur >= len(candidates) || total > target {
        return
    }

    // to include in sum
    temp = append(temp, candidates[cur])
    backtrack(candidates, target, res, temp, cur, total+candidates[cur])

    // to not include
    temp = temp[:len(temp)-1]
    backtrack(candidates, target, res, temp, cur+1, total)
}

func main() {
    candidates := []int{2, 3, 6, 7}
    target := 7
    fmt.Println(combinationSum(candidates, target))  // Output: [[2 2 3] [7]]
}

/*
Here's a breakdown of the Go code:

- `combinationSum(candidates []int, target int) [][]int`: This is the main function that takes a slice of integers (`candidates`) and an integer (`target`) as input and returns a slice of slices of integers. The returned slice contains all combinations of `candidates` that sum up to `target`.

- `res := [][]int{}`: This line initializes `res`, which is a slice that will store all valid combinations.

- `backtrack(candidates, target, &res, []int{}, 0, 0)`: This line is the initial call to the `backtrack` function. It passes `candidates`, `target`, a reference to `res`, an empty slice for the current combination (`temp`), and `0` for the current index (`cur`) and the current sum (`total`).

- `func backtrack(candidates []int, target int, res *[][]int, temp []int, cur int, total int)`: This is the declaration of the `backtrack` function. It's a recursive function that generates all combinations that sum up to `target`.

- `if total == target {...}`: This `if` statement checks if `total` equals `target`. If it does, it adds a copy of `temp` to `res` and returns.

- `if cur >= len(candidates) || total > target {...}`: This `if` statement checks if `cur` has reached the end of `candidates` or if `total` exceeds `target`. If either condition is true, it returns without making any changes to `res`.

- `temp = append(temp, candidates[cur])` and `backtrack(candidates, target, res, temp, cur, total+candidates[cur])`: These lines include the current number in the combination and the sum, and make a recursive call to `backtrack`.

- `temp = temp[:len(temp)-1]` and `backtrack(candidates, target, res, temp, cur+1, total)`: These lines exclude the current number from the combination and the sum, and make a recursive call to `backtrack`.

- `func main() {...}`: This is the main function that runs when the program starts. It initializes `candidates` and `target`, and prints the result of `combinationSum(candidates, target)`.

*/