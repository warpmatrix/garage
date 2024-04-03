#include <algorithm>
#include <iostream>
#include <vector>

class Solution {
   public:
    int minimumAddedCoins(std::vector<int>& coins, int target) {
        std::sort(coins.begin(), coins.end());
        int subTarget = 1, idx = 0, ret = 0;
        while (subTarget <= target) {
            if (idx < coins.size() && coins[idx] <= subTarget) {
                subTarget += coins[idx];
                idx++;
            } else {
                ret++;
                subTarget *= 2;
            }
        }
        return ret;
    }
};

int main(int argc, char const* argv[]) {
    auto s = Solution();
    auto coins = std::vector<int>({1, 4, 10});
    s.minimumAddedCoins(coins, 19);
    return 0;
}
