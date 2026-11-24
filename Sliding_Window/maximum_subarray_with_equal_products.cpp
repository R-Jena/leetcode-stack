/*
{
    "problem_name": "Maximum Subarray With Equal Products",
    "category": "Sliding Window / Bit Manipulation & Number Theory",
    "time_complexity": "O(N) -> Single sliding window pass using 4-bit prime factor masks. Zero GCD/LCM calls.",
    "space_complexity": "O(1) -> Constant 11-byte lookup table and integer registers. Zero heap allocations."
}
*/
#include <vector>
#include <algorithm>

using namespace std;

class Solution {
public:
    int maxLength(vector<int>& nums) {
        // Prime factor bitmasks for numbers 0..10 using primes {2, 3, 5, 7}
        // bit 0 -> 2, bit 1 -> 3, bit 2 -> 5, bit 3 -> 7
        static constexpr int mask[11] = {
            0,      // 0 (unused)
            0b0000, // 1
            0b0001, // 2
            0b0010, // 3
            0b0001, // 4 (2^2)
            0b0100, // 5
            0b0011, // 6 (2 * 3)
            0b1000, // 7
            0b0001, // 8 (2^3)
            0b0010, // 9 (3^2)
            0b0101  // 10 (2 * 5)
        };

        // Any subarray of length 2 always satisfies a * b == gcd(a, b) * lcm(a, b)
        int ans = 2;
        int windowMask = 0;
        int left = 0;
        int n = nums.size();

        // For length >= 3, prod == gcd * lcm holds iff all elements are pairwise coprime
        for (int right = 0; right < n; ++right) {
            int m = mask[nums[right]];

            // Shrink window from the left until no prime factor conflicts remain
            while (windowMask & m) {
                windowMask ^= mask[nums[left]];
                ++left;
            }

            windowMask |= m;
            ans = max(ans, right - left + 1);
        }

        return ans;
    }
};
