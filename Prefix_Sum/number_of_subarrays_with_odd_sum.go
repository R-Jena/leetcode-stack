/*
	{
	    "problem_name": "Number of Subarrays With Odd Sum",
	    "category": "Prefix Sum / Bit Manipulation",
	    "time_complexity": "O(N) -> Pure linear scan, bitwise parity tracking.",
	    "space_complexity": "O(1) -> Only 4 CPU registers. Absolute zero allocations."
	}
*/
package main

func numOfSubarrays(arr []int) int {
	var odd, even int = 0, 1
	var ans int = 0
	var parity int = 0

	for _, num := range arr {
		// HFT Trick 1: Bitwise XOR to toggle parity instead of modulo.
		// If num is odd (lowest bit 1), parity flips. If even (0), parity stays.
		parity ^= num & 1

		if parity == 1 {
			// Current prefix sum is odd.
			// It forms an odd subarray with every previous EVEN prefix sum.
			ans += even
			odd++
		} else {
			// Current prefix sum is even.
			// It forms an odd subarray with every previous ODD prefix sum.
			ans += odd
			even++
		}
	}

	// HFT Trick 2: Execute modulo ONLY ONCE at the absolute end.
	// Max answer is ~5*10^9, easily fitting inside a 64-bit int.
	// This skips 100,000 expensive modulo operations inside the loop!
	return ans % 1000000007
}
