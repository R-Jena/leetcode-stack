/*
	{
	    "problem_name": "Minimum Positive Sum Subarray",
	    "category": "Arrays / Micro-Optimized Prefix Sum",
	    "time_complexity": "O(N^2) -> Strictly bounded nested loops. Zero wasted iterations.",
	    "space_complexity": "O(1) -> 105-element array allocated entirely on the CPU stack (zero GC heap allocation)."
	}
*/
package main

func minimumSumSubarray(nums []int, l int, r int) int {
	// HFT Trick 1: Stack Allocation!
	// Limit is N=100. Instead of make([]int, n), we hardcode [105]int.
	// Go compiles this directly into the stack frame. Zero garbage collection!
	var P [105]int
	n := len(nums)

	// Build prefix sum array in O(N)
	for i := 0; i < n; i++ {
		P[i+1] = P[i] + nums[i]
	}

	// HFT Trick 2: Infinity initialization.
	// Max possible sum is 100 * 1000 = 100,000.
	// This completely eliminates the "if minSum == -1" branch from the hot loop.
	minSum := 100000000

	for i := 0; i <= n-l; i++ {
		// HFT Trick 3: Pre-calculate loop bounds outside.
		// Inner loop only iterates exactly over valid lengths (l to r).
		// No internal "if length > r { break }" needed.
		end := i + r
		if end > n {
			end = n
		}

		for j := i + l; j <= end; j++ {
			sum := P[j] - P[i]
			// Single clean condition. Supreme branch prediction.
			if sum > 0 && sum < minSum {
				minSum = sum
			}
		}
	}

	if minSum == 100000000 {
		return -1
	}
	return minSum
}
