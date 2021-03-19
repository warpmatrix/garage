#include <iostream>
#include <queue>
#include <vector>

std::vector<int> maxSlidingWindow(std::vector<int>& nums, int k) {
    std::deque<int> deque;
    std::vector<int> res;
    for (size_t i = 0; i < nums.size(); i++) {
        while (deque.size() && nums[deque.back()] <= nums[i]) deque.pop_back();
        deque.push_back(i);
        if (deque.front() + k <= i) deque.pop_front();
        if (i + 1 >= k) res.push_back(nums[deque.front()]);
    }
    return res;
}
