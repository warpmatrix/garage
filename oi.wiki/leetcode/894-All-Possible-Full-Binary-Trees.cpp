#include <cstddef>
#include <vector>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right)
        : val(x), left(left), right(right) {}
};

class Solution {
   public:
    std::vector<TreeNode *> allPossibleFBT(int n) {
        std::vector<TreeNode *> ret;
        if (n % 2 == 0) {
            return ret;
        }
        if (n == 1) {
            ret.push_back(new TreeNode());
            return ret;
        }
        for (size_t i = 1; i < n; i += 2) {
            auto leftTrees = allPossibleFBT(i);
            auto rightTrees = allPossibleFBT(n - i - 1);
            for (auto &&leftTree : leftTrees) {
                for (auto &&rightTree : rightTrees) {
                    auto root = new TreeNode(0, leftTree, rightTree);
                    ret.push_back(root);
                }
            }
        }
        return ret;
    }
};
