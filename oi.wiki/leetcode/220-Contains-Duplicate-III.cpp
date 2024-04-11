#include <climits>
#include <set>
#include <vector>
using std::vector;

bool containsNearbyAlmostDuplicate(vector<int>& nums, int k, int t) {
    std::set<int> s;
    for (int i = 0; i < nums.size(); i++) {
        // max(num[i] - t, INT_MIN)
        auto iter = s.lower_bound(std::max(nums[i], INT_MIN + t) - t);
        // min(nums[i] + t, INT_MAX)
        if (iter != s.end() && *iter <= std::min(nums[i], INT_MAX - t) + t) {
            return true;
        }
        s.insert(nums[i]);
        if (s.size() > k) {
            s.erase(nums[i] - k);
        }
    }
    return false;
}