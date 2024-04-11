#include <deque>
#include <iostream>
#include <string>

class Solution {
   public:
    std::string finalString(std::string s) {
        std::deque<char> q;
        bool frontIsHead = true;
        for (auto &&ch : s) {
            if (ch == 'i') {
                frontIsHead = !frontIsHead;
                continue;
            }
            if (frontIsHead) {
                q.push_back(ch);
            } else {
                q.push_front(ch);
            }
        }
        if (!frontIsHead) {
            return std::string(q.crbegin(), q.crend());
        }
        return std::string(q.cbegin(), q.cend());
    }
};

int main(int argc, char const *argv[]) {
    auto s = Solution();
    std::cout << s.finalString("string") << '\n';
    std::cout << s.finalString("poiinter") << '\n';
    std::cout << s.finalString("stringing") << '\n';
    return 0;
}
