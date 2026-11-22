/*
{
    "problem_name": "Check if All A's Appears Before All B's",
    "category": "Strings / Single-Pass Scan",
    "time_complexity": "O(N) -> Single linear pass with immediate early exit on 'ba' transition.",
    "space_complexity": "O(1) -> Pure CPU register loop. Absolute zero heap allocations."
}
*/
#include <string>

using namespace std;

class Solution {
public:
    bool checkString(string s) {
        int n = s.size();
        // Every 'a' appears before every 'b' iff the "ba" transition never occurs
        for (int i = 1; i < n; ++i) {
            if (s[i - 1] == 'b' && s[i] == 'a') {
                return false;
            }
        }
        return true;
    }
};
