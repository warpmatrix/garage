#include <string>
#include <vector>
using std::string;

int numDistinct(string s, string t) {
    if (s.length() < t.length()) return 0;
    std::vector<unsigned int> locs[52];
    for (int i = 0; i < s.length(); i++) {
        if (s[i] >= 'a' && s[i] <= 'z') s[i] = s[i] - 'a' + 'A' + 26;
    }
    for (int i = 0; i < t.length(); i++) {
        if (t[i] >= 'a' && t[i] <= 'z') t[i] = t[i] - 'a' + 'A' + 26;
    }
    for (int i = 0; i < s.length(); i++) {
        locs[s[i] - 'A'].push_back(i);
    }
    std::vector<std::vector<unsigned int>> dp(t.length(), std::vector<unsigned int>());
    for (int i = t.length() - 1; i >= 0; i--) {
        dp[i].resize(locs[t[i] - 'A'].size());
        for (int j = 0; j < locs[t[i] - 'A'].size(); j++) {
            if (i == t.length() - 1) {
                dp[i][j] = 1;
                continue;
            }
            for (int k = locs[t[i + 1] - 'A'].size() - 1; k >= 0; k--) {
                if (locs[t[i] - 'A'][j] >= locs[t[i + 1] - 'A'][k]) break;
                dp[i][j] += dp[i + 1][k];
            }
        }
    }
    unsigned int ans = 0;
    for (int i = 0; i < dp[0].size(); i++) {
        ans += dp[0][i];
    }
    return ans;
}

// #include <memory.h>

// int numDistinct(string s, string t) {
//     if (s.length() < t.length()) return 0;
//     long dp[s.length() + 1][t.length() + 1];
//     memset(dp, 0, sizeof(dp));
//     for (int i = 0; i <= s.length(); i++) {
//         dp[i][t.length()] = 1;
//     }
//     for (int i = s.length() - 1; i >= 0; i--) {
//         for (int j = t.length() - 1; j >= 0; j--) {
//             if (s[i] == t[j]) {
//                 dp[i][j] = dp[i + 1][j + 1] + dp[i + 1][j];
//             } else {
//                 dp[i][j] = dp[i + 1][j];
//             }
//         }
//     }
//     return dp[0][0];
// }
