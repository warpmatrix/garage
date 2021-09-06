#include <algorithm>
#include <vector>

class Solution {
public:
    int search(std::vector<int>& nums, int target) {
        auto iter = std::lower_bound(nums.begin(), nums.end(), target);
        if (iter == nums.end() || target < *iter) { return -1; }
        return iter - nums.begin();
    }
};
