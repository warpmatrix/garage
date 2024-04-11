#include <stack>
#include <vector>
using namespace std;

vector<int> nextGreaterElements(vector<int>& nums) {
    if (nums.empty()) return vector<int>();
    std::vector<int> ret(nums.size(), -1);
    std::stack<int> stack;
    for (int i = 0; i < nums.size(); i++) {
        while (!stack.empty()) {
            int idx = stack.top();
            if (nums[i] <= nums[idx]) break;
            ret[idx] = nums[i];
            stack.pop();
        }
        stack.push(i);
    }
    for (int i = 0; i < nums.size(); i++) {
        int idx = stack.top();
        while (nums[i] > nums[idx]) {
            ret[idx] = nums[i];
            stack.pop();
            idx = stack.top();
        }
    }
    return ret;
}