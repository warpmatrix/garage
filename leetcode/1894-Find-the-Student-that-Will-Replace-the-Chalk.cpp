#include <vector>
#include <algorithm>

using std::vector;

class Solution {
public:
    int chalkReplacer(vector<int>& chalk, int k) {
        for (int i = 0; i < chalk.size() - 1; i++) {
            if (chalk[i] > k) return i;
            chalk[i + 1] += chalk[i];
        }
        k %= chalk[chalk.size() - 1];
        auto iter = std::upper_bound(chalk.begin(), chalk.end(), k);
        return iter - chalk.begin();
    }
};
