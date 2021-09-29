#include <vector>
using std::vector;

class Solution {
public:
    int findPeakElement(vector<int>& nums) {
        auto lt = [&](int pos) -> bool {
            return pos == 0 || nums[pos - 1] < nums[pos];
        };
        auto gt = [&](int pos) -> bool {
            return pos == nums.size() - 1 || nums[pos] > nums[pos + 1];
        };
        auto isPeak = [&](int pos) -> bool {
            return lt(pos) && gt(pos);
        };
        int lptr = 0, rptr = nums.size() - 1;
        int ret = (lptr + rptr) / 2;
        while (!isPeak(ret)) {
            if (lt(ret)) {
                lptr = ret + 1;
            } else {
                rptr = ret - 1;
            }
            ret = (lptr + rptr) / 2;
        }
        return ret;
    }
};
