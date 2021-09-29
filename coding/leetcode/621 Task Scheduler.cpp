#include <memory.h>

#include <iostream>
#include <vector>
using namespace std;

int leastInterval(vector<char>& tasks, int n) {
    int cnts[26];
    memset(cnts, 0, sizeof(cnts));
    int maxCnt = 0, maxIdx = -1, maxNum = 0;
    for (auto&& task : tasks) {
        int index = task - 'A';
        cnts[index]++;
        if (cnts[index] == maxCnt)
            maxNum++;
        else if (cnts[index] > maxCnt) {
            maxCnt = cnts[index];
            maxIdx = index;
            maxNum = 1;
        }
    }
    int ret = (maxCnt - 1) * (n + 1) + maxNum;
    if (ret < tasks.size()) ret = tasks.size();
    return ret;
}

int main(int argc, char const* argv[]) {
    vector<char> tasks;
    tasks = vector<char>({'A', 'A', 'A', 'B', 'B', 'B'});
    std::cout << leastInterval(tasks, 2) << '\n';
    std::cout << leastInterval(tasks, 0) << '\n';
    tasks = vector<char>(
        {'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'});
    std::cout << leastInterval(tasks, 2) << '\n';
    return 0;
}
