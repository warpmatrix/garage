#include <algorithm>
#include <iostream>
#include <string>

class Solution {
   public:
    std::string maximumBinaryString(std::string binary) {
        int zeroPos = binary.find('0');
        if (zeroPos == std::string::npos) {
            return binary;
        }
        int zeroCnt = std::count(binary.begin(), binary.end(), '0');
        std::string ret(binary.size(), '1');
        ret[zeroPos + zeroCnt - 1] = '0';
        return ret;
    }
};

int main(int argc, char const *argv[]) {
    auto s = Solution();
    std::cout << s.maximumBinaryString("000110") << '\n';
    std::cout << s.maximumBinaryString("01") << '\n';
    std::cout << s.maximumBinaryString("10") << '\n';
    return 0;
}
