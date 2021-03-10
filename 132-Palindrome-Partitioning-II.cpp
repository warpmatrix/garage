#include <algorithm>
#include <string>
#include <vector>
using std::string;
using std::vector;

bool isCheck[2000][2000] = {};
bool isPalindrome[2000][2000];

bool checkPalindrome(int head, int tail, const string &s) {
    if (isCheck[head][tail]) return isPalindrome[head][tail];
    if (head + 2 <= tail) {
        isPalindrome[head][tail] =
            s[head] == s[tail] && checkPalindrome(head + 1, tail - 1, s);
    } else {
        isPalindrome[head][tail] = true;
    }
    isCheck[head][tail] = true;
    return isPalindrome[head][tail];
}

int minCut(string s) {
    int minCut[s.length()];
    minCut[0] = 0;
    for (int i = 1; i < s.length(); i++) {
        if (checkPalindrome(0, i, s) == true) {
            minCut[i] = 0;
            continue;
        }
        minCut[i] = 0x3f3f3f3f;
        for (int j = 0; j < i; j++) {
            if (checkPalindrome(j + 1, i, s) == true)
                minCut[i] = std::min(minCut[i], minCut[j] + 1);
        }
    }
    return minCut[s.length() - 1];
}
