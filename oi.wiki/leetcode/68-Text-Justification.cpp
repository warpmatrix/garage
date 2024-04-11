#include <string>
#include <vector>

using std::string;
using std::vector;

vector<string> fullJustify(vector<string>& words, int maxWidth) {
    std::vector<std::string> ret;
    int len = words[0].size(), idx = 0;
    for (size_t i = 1; i < words.size(); i++) {
        // including the length of each word's space
        // except for the first word in every line
        int nLen = len + words[i].size() + 1;
        if (nLen <= maxWidth) {
            len = nLen;
        } else {
            int cnt = i - idx;
            int wordsLen = len - cnt + 1;
            ret.push_back(
                constructLine(words.begin() + idx, cnt, maxWidth, wordsLen));
            len = words[i].size();
            idx = i;
        }
    }
    std::string lastLine(words[idx]);
    for (int i = idx + 1; i < words.size(); i++) {
        lastLine += " " + words[i];
    }
    lastLine += std::string(maxWidth - len, ' ');
    ret.push_back(lastLine);
    return ret;
}

std::string constructLine(std::vector<std::string>::const_iterator begin,
                          int wordCnt, int maxWidth, int wordsLen) {
    if (wordCnt == 1) {
        return *begin + std::string(maxWidth - wordsLen, ' ');
    }
    int spaceLen = (maxWidth - wordsLen) / (wordCnt - 1);
    std::string line(*begin);
    for (size_t j = 1; j < wordCnt; j++) {
        if (j - 1 < (maxWidth - wordsLen) % (wordCnt - 1)) {
            line += std::string(spaceLen + 1, ' ');
        } else {
            line += std::string(spaceLen, ' ');
        }
        line += *(begin + j);
    }
    return line;
}
