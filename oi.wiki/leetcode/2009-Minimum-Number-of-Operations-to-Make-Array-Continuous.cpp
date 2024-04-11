#include <vector>
#include <algorithm>

class Solution {
public:
    int minOperations(std::vector<int>& nums) {
        int ret = nums.size();
        std::sort(nums.begin(), nums.end());
        auto last = std::unique(nums.begin(), nums.end());
        auto right = nums.begin();
        for (auto left = nums.begin(); left < last; left++) {
            int right_val = *left + nums.size() - 1;
            while (right < last && *right <= right_val) {
                right++;
            }
            int cnt = nums.size() - (right - left);
            ret = std::min(ret, cnt);
        }
        return ret;
    }
};
