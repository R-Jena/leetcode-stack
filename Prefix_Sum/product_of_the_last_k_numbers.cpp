/*
{
    "problem_name": "Product of the Last K Numbers",
    "category": "Prefix Sum / Multiplicative Prefix Array",
    "time_complexity": "O(1) -> Amortized constant time for add() and strict O(1) for getProduct().",
    "space_complexity": "O(N) -> Stores cumulative non-zero prefix products since the last zero."
}
*/
#include <vector>

using namespace std;

class ProductOfNumbers {
public:
    vector<int> prefix;

    ProductOfNumbers() {
        // Base identity element (1) so prefix multiplication and division stay valid
        prefix.push_back(1);
    }

    void add(int num) {
        if (num == 0) {
            // A zero invalidates all previous products; reset the prefix stream
            prefix.clear();
            prefix.push_back(1);
        } else {
            // Append running cumulative product
            prefix.push_back(prefix.back() * num);
        }
    }

    int getProduct(int k) {
        // If k reaches or crosses the last zero reset boundary, the product is strictly 0
        if (k >= (int)prefix.size()) return 0;

        // Divide total cumulative product by the product right before the last k elements
        return prefix.back() / prefix[prefix.size() - 1 - k];
    }
};
