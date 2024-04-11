#include <algorithm>
#include <string>
#include <vector>

using std::string;
using std::vector;

string findLongestWord(string s, vector<string>& dictionary) {
    std::sort(dictionary.begin(), dictionary.end());
    std::vector<int> idxs(dictionary.size(), 0);
    for (auto&& ch : s) {
        for (int i = 0; i < dictionary.size(); i++) {
            if (idxs[i] < dictionary[i].length() &&
                ch == dictionary[i][idxs[i]]) {
                idxs[i]++;
            }
        }
    }
    int max = 0, retIdx = -1;
    for (int i = 0; i < dictionary.size(); i++) {
        if (idxs[i] == dictionary[i].length() && max < idxs[i]) {
            max = idxs[i];
            retIdx = i;
        }
    }
	return (retIdx != -1) ? dictionary[retIdx] : "";
}