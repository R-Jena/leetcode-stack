/*
{
    "problem_name": "Stone Removal Game",
    "category": "Math / Game Theory Simulation",
    "time_complexity": "O(1) -> Bounded to at most 7 iterations since 1 <= n <= 50.",
    "space_complexity": "O(1) -> Strictly 2 CPU registers (take, aliceTurn). Zero heap allocations."
}
*/
class Solution {
public:
    bool canAliceWin(int n) {
        int take = 10;
        bool aliceTurn = true;

        // Simulate turns until the current player cannot make a valid move
        while (n >= take) {
            n -= take;
            --take;
            aliceTurn = !aliceTurn;
        }

        // If the loop exits on Bob's turn (aliceTurn == false), Bob loses and Alice wins
        return !aliceTurn;
    }
};
