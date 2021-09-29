#include <vector>
#include <string>
using namespace std;

vector<string> summaryRanges(vector<int>& nums) {
    vector<string> ret;
    int len = nums.size();
    int pos = 0;
    while (pos < len) {
        int lowBound = pos;
        for (pos = pos + 1; pos < len && nums[pos] == nums[pos - 1] + 1; pos++)
            ;
        int highBound = pos - 1;
        string tmp = to_string(nums[lowBound]);
        if (lowBound < highBound) {
            tmp.append("->");
            tmp.append(to_string(nums[highBound]));
        }
        ret.push_back(move(tmp));
    }
    return ret;
}
