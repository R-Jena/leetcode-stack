/*
	{
	    "problem_name": "Minimum Score by Changing Two Elements",
	    "category": "Greedy / Single-Pass Extremes",
	    "time_complexity": "O(N) -> Bypasses O(N log N) sorting by tracking top-3 and bottom-3 in one linear scan.",
	    "space_complexity": "O(1) -> Strictly 6 integer registers. Absolute zero heap allocations."
	}
*/
package main

func minimizeSum(nums []int) int {
	// Constraints: 1 <= nums[i] <= 10^9
	const INF = 1_000_000_001
	min1, min2, min3 := INF, INF, INF
	max1, max2, max3 := 0, 0, 0

	// Single bare-metal pass to extract the 3 smallest and 3 largest elements
	for _, x := range nums {
		if x < min1 {
			min3, min2, min1 = min2, min1, x
		} else if x < min2 {
			min3, min2 = min2, x
		} else if x < min3 {
			min3 = x
		}

		if x > max1 {
			max3, max2, max1 = max2, max1, x
		} else if x > max2 {
			max3, max2 = max2, x
		} else if x > max3 {
			max3 = x
		}
	}

	// Low score is always 0 by matching changed elements to existing ones.
	// Evaluate the 3 optimal high-score candidates:
	ans := max1 - min3 // Change 2 smallest
	if d := max3 - min1; d < ans {
		ans = d // Change 2 largest
	}
	if d := max2 - min2; d < ans {
		ans = d // Change 1 smallest and 1 largest
	}

	return ans
}
