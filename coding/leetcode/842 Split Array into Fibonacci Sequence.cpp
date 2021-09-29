#include <climits>
#include <iostream>
#include <string>
#include <vector>
using namespace std;

bool isValid(int first, int second, string S, vector<int> &list) {
    vector<int> tmp = list;
    int index = 0, len = S.length();
    while (index < len) {
        long long int sum = first + second;
        string sumStr = to_string(sum);
        int sumLen = sumStr.length();
        if (sum >= INT32_MAX) return false;
        if (sumStr == S.substr(index, sumLen)) {
            index += sumLen;
            tmp.push_back(sum);
            first = second;
            second = sum;
        } else {
            return false;
        }
    }
    list = tmp;
    return true;
}

vector<int> splitIntoFibonacci(string S) {
    vector<int> ret;
    int len = S.length();
    for (int fstLen = 1; fstLen < len - 2; fstLen++) {
        long long int fstNum = stoll(S.substr(0, fstLen));
        if (fstNum >= INT32_MAX) break;
        ret.push_back(fstNum);
        for (int secLen = 1; secLen < len - fstLen - 1; secLen++) {
            long long int secNum = stoll(S.substr(fstLen, secLen));
            if (secNum >= INT32_MAX) break;
            ret.push_back(secNum);
            if (isValid(fstNum, secNum, S.substr(fstLen + secLen), ret)) return ret;
            ret.pop_back();
        }
        ret.pop_back();
    }
    return ret;
}

int main(int argc, char const *argv[]) {
    vector<int> v = splitIntoFibonacci("123456579");
    for (auto &&num : v) {
        std::cout << num << ' ';
    }
    std::cout << '\n';
    return 0;
}
