#include <stack>
#include <string>
using std::string;

string removeDuplicates(string S) {
    std::string ret;
    for (int i = 0; i < S.length(); i++) {
        if (ret.empty() || S[i] != ret.back())
            ret.push_back(S[i]);
        else
            ret.pop_back();
    }
    return ret;
}
