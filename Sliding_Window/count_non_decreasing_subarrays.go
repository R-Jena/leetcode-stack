/*
	{
	    "problem_name": "Number of Subarrays That Can Be Made Non-Decreasing",
	    "category": "Sliding Window / Monotonic Deque",
	    "time_complexity": "O(N) -> Pure linear scan. Each element pushed/popped at most once.",
	    "space_complexity": "O(N) -> Single pre-allocated slice acting as a ring-free Deque. Absolute zero runtime allocations."
	}
*/
package main

// Changed 'k' parameter back to 'int' to match LeetCode's exact hidden signature
func countNonDecreasingSubarrays(nums []int, k int) int64 {
	n := len(nums)

	// Pre-allocate a single slice to act as a double-ended queue (Deque)
	q := make([]int, n+1)
	head, tail := n, n

	var cost int64 = 0
	var ans int64 = 0
	R := n - 1
	k64 := int64(k) // Cast to int64 internally for overflow-safe comparisons

	for L := n - 1; L >= 0; L-- {
		// 1. Prepend nums[L] and maintain monotonic decreasing property
		for head < tail && nums[L] >= nums[q[head]] {
			popped := q[head]
			head++

			var length int
			if head < tail {
				length = q[head] - popped
			} else {
				length = R + 1 - popped
			}

			cost += int64(nums[L]-nums[popped]) * int64(length)
		}

		// 2. Push current index to the front of the deque
		head--
		q[head] = L

		// 3. Shrink window from the right if the cost exceeds k64
		for cost > k64 {
			back := q[tail-1]

			cost -= int64(nums[back] - nums[R])
			R--

			if tail > head && q[tail-1] > R {
				tail--
			}
		}

		// 4. Add all valid subarrays ending between L and R
		ans += int64(R - L + 1)
	}

	return ans
}
