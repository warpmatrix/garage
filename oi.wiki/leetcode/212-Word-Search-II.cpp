#include <algorithm>
#include <set>
#include <string>
#include <vector>
using std::string;
using std::vector;

class TrieNode {
   public:
    std::vector<TrieNode *> children;
    std::string word;
    TrieNode() : children(26), word("") {}
    void insert(std::string &word) {
        TrieNode *node = this;
        for (auto &&ch : word) {
            if (!node->children[ch - 'a']) {
                node->children[ch - 'a'] = new TrieNode();
            }
            node = node->children[ch - 'a'];
        }
        node->word = word;
    }
};

typedef std::pair<int, int> Loc;

class Solution {
   public:
    vector<string> findWords(vector<vector<char>> &board,
                             vector<string> &words) {
        TrieNode *trie = new TrieNode();
        for (auto &&word : words) {
            trie->insert(word);
        }
        std::set<std::string> ret;
        const std::vector<Loc> dirs{{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        auto dfs = [&](auto &&self, Loc loc, TrieNode *node) -> void {
            if (board[loc.first][loc.second] == '#') return;
            char ch = board[loc.first][loc.second];
            if (!node->children[ch - 'a']) return;
            node = node->children[ch - 'a'];
            if (node->word.length() > 0) {
                ret.insert(node->word);
            }
            board[loc.first][loc.second] = '#';
            for (auto &&dir : dirs) {
                Loc nloc = {loc.first + dir.first, loc.second + dir.second};
                if (!(nloc.first >= 0 && nloc.first < board.size() &&
                      nloc.second >= 0 && nloc.second < board[0].size()))
                    continue;
                self(self, nloc, node);
            }
            board[loc.first][loc.second] = ch;
        };
        for (int i = 0; i < board.size(); i++) {
            for (int j = 0; j < board[0].size(); j++) {
                dfs(dfs, {i, j}, trie);
            }
        }
        return vector<std::string>(ret.begin(), ret.end());
    }
};
