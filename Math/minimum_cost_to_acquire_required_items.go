/*
	{
	    "problem_name": "Minimum Total Cost of Items",
	    "category": "Math / Greedy Optimization",
	    "time_complexity": "O(1) -> Absolute bare-metal mathematical calculation. Zero loops.",
	    "space_complexity": "O(1) -> Zero heap allocations, purely 64-bit CPU registers."
	}
*/
package main

func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	var totalCost int64 = 0

	// Case 1: Separate items are cheaper.
	// Completely ignore the combo (Type 3) and fulfill requirements purely individually.
	if cost1+cost2 <= costBoth {
		totalCost = int64(cost1)*int64(need1) + int64(cost2)*int64(need2)
	} else {
		// Case 2: The combo pack (Type 3) is a steal!
		// Maximize its use until at least one requirement is completely fulfilled.
		cnt3 := min(need1, need2)
		totalCost = int64(cnt3) * int64(costBoth)

		// Deduct the overlapping fulfilled quantities
		need1 -= cnt3
		need2 -= cnt3

		// Residual Execution: Only one of these will be > 0.
		// We use min(costX, costBoth) because in extremely skewed test cases,
		// buying an extra combo pack might STILL be cheaper than buying a single item!
		if need1 > 0 {
			totalCost += int64(need1) * int64(min(cost1, costBoth))
		} else if need2 > 0 {
			totalCost += int64(need2) * int64(min(cost2, costBoth))
		}
	}

	return totalCost
}
