#include <algorithm>

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
    int maxDiffFromTop(TreeNode *node, int min, int max) {
        int ret = max - min;
        if (node->left) {
            int newMin = std::min(min, node->left->val);
            int newMax = std::max(max, node->left->val);
            int leftRet = maxDiffFromTop(node->left, newMin, newMax);
            ret = std::max(ret, leftRet);
        }
        if (node->right) {
            int newMin = std::min(min, node->right->val);
            int newMax = std::max(max, node->right->val);
            int rightRet = maxDiffFromTop(node->right, newMin, newMax);
            ret = std::max(ret, rightRet);
        }
        return ret;
    }

    int maxAncestorDiff(TreeNode *root) {
        int ret = maxDiffFromTop(root, root->val, root->val);
        return ret;
    }
};
