/*
	{
	    "problem_name": "Longest Strictly Increasing or Strictly Decreasing Subarray",
	    "category": "Arrays / State Machine",
	    "time_complexity": "O(N) -> Single bare-metal linear scan.",
	    "space_complexity": "O(1) -> Only 3 CPU registers used (maxLen, inc, dec). Absolute zero allocations."
	}
*/
package main

func longestMonotonicSubarray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	maxLen, inc, dec := 1, 1, 1

	// Hardware-friendly single pass loop
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			inc++
			dec = 1
		} else if nums[i] < nums[i-1] {
			dec++
			inc = 1
		} else {
			inc = 1
			dec = 1
		}

		// Branchless-style max updates
		if inc > maxLen {
			maxLen = inc
		}
		if dec > maxLen {
			maxLen = dec
		}
	}

	return maxLen
}
