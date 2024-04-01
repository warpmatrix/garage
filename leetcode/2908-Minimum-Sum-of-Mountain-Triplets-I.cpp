#include <climits>
#include <cstddef>
#include <vector>

class Solution {
   public:
    int minimumSum(std::vector<int> &nums) {
        size_t len = nums.size();
        std::vector<int> leftMin(len);
        leftMin[0] = nums[0];
        for (size_t i = 1; i < nums.size(); i++) {
            leftMin[i] = std::min(leftMin[i - 1], nums[i]);
        }
        std::vector<int> rightMin(len);
        rightMin[len - 1] = nums[len - 1];
        for (int i = len - 2; i >= 0; i--) {
            rightMin[i] = std::min(rightMin[i + 1], nums[i]);
        }
        int ret = INT_MAX;
        for (size_t i = 1; i < len - 1; i++) {
            if (leftMin[i - 1] < nums[i] && nums[i] > rightMin[i + 1]) {
                ret = std::min(ret, leftMin[i - 1] + nums[i] + rightMin[i + 1]);
            }
        }
        if (ret == INT_MAX) {
            ret = -1;
        }
        return ret;
    }
};

int main(int argc, char const *argv[]) {
    auto s = Solution();
    return 0;
}
