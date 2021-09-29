#include <string>
#include <vector>
using std::string;
using std::vector;

bool isCheck[16][16] = {false};
bool isPalindrome[16][16];

bool checkPalindrome(int head, int tail, const string &s) {
    if (isCheck[head][tail]) return isPalindrome[head][tail];
    if (head >= tail) {
        isPalindrome[head][tail] = true;
    } else {
        isPalindrome[head][tail] =
            s[head] == s[tail] && checkPalindrome(head + 1, tail - 1, s);
    }
    isCheck[head][tail] = true;
    return isPalindrome[head][tail];
}

void dfs(int head, const string &s, vector<vector<string>> &ret,
         vector<string> &subStrs) {
    if (head == s.length()) {
        ret.push_back(subStrs);
        return;
    }
    for (int tail = head; tail < s.length(); tail++) {
        if (checkPalindrome(head, tail, s) == true) {
            subStrs.push_back(s.substr(head, tail - head + 1));
            dfs(tail + 1, s, ret, subStrs);
            subStrs.pop_back();
        }
    }
}

vector<vector<string>> partition(string s) {
    vector<vector<string>> ret;
    vector<string> subStrs;
    dfs(0, s, ret, subStrs);
    return ret;
}