/*
	{
	    "problem_name": "Maximize Consecutive Elements in an Array After Modification",
	    "category": "Dynamic Programming / Rolling State Machine",
	    "time_complexity": "O(N log N) -> Zero-boxing slices.Sort + single linear state-machine pass.",
	    "space_complexity": "O(1) -> Pure 3-register rolling window (a, b, c). Absolute zero heap allocations."
	}
*/
package main

import "slices"

func maxSelectedElements(nums []int) int {
	// Sort in ascending order so we can process consecutive chains sequentially
	slices.Sort(nums)

	// Rolling DP state represented by 3 CPU registers relative to the current element x:
	// a -> dp[x-1] (longest consecutive sequence ending at x-1)
	// b -> dp[x]   (longest consecutive sequence ending at x)
	// c -> dp[x+1] (longest consecutive sequence ending at x+1)
	a, b, c := 0, 0, 0
	prev, ans := nums[0], 0

	for _, x := range nums {
		// Shift the rolling window based on the gap between current element x and previous element
		switch x - prev {
		case 0:
			// Duplicate element: window stays aligned at the exact same values (x-1, x, x+1)
		case 1:
			// Shifted right by 1: old dp[x] becomes dp[x-1], old dp[x+1] becomes dp[x]
			a, b, c = b, c, 0
		case 2:
			// Shifted right by 2: old dp[x+1] becomes new dp[x-1], others reset
			a, b, c = c, 0, 0
		default:
			// Gap > 2: previous consecutive chains are completely broken, reset all states
			a, b, c = 0, 0, 0
		}

		prev = x

		// Transition 1: Increment x to x+1 -> extends chain ending at x (b)
		// MUST be evaluated before updating b to prevent reusing the same element twice
		c = max(c, b+1)

		// Transition 2: Keep x unchanged -> extends chain ending at x-1 (a)
		b = max(b, a+1)

		// Track global maximum consecutive length
		ans = max(ans, b, c)
	}

	return ans
}
