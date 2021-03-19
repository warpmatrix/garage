#include <iostream>
#include <string>
#include <vector>
using namespace std;

string reorganizeString(string S) {
    std::vector<int> cnt(26, 0);
    int maxCnt = 0;
    for (auto &&ch : S) {
        cnt[ch - 'a']++;
        maxCnt = std::max(maxCnt, cnt[ch - 'a']);
    }
    // can't reorganize
    if (maxCnt > (S.size() + 1) / 2) return "";
    // reorganize string
    string ret(S.size(), '\0');
    int evenIdx = 0, oddIdx = 1;
    for (int i = 0; i < 26; i++) {
        while (cnt[i] && cnt[i] <= S.size() / 2 && oddIdx < S.size()) {
            ret[oddIdx] = 'a' + i;
            oddIdx += 2;
            cnt[i]--;
        }
        while (cnt[i]) {
            ret[evenIdx] = 'a' + i;
            evenIdx += 2;
            cnt[i]--;
        }
    }
    return ret;
}

int main(int argc, char const *argv[]) {
    string str = "aba";
    std::cout << reorganizeString(str) << '\n';
    return 0;
}
