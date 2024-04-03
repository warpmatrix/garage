#include <queue>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

class Solution {
public:
    TreeNode* getTargetCopy(TreeNode* original, TreeNode* cloned, TreeNode* target) {
        std::queue<TreeNode *> originQ({original});
        std::queue<TreeNode *> clonedQ({cloned});
        while (originQ.size()) {
            TreeNode *originNode = originQ.front();
            TreeNode *clonedNode = clonedQ.front();
            originQ.pop(), clonedQ.pop();
            if (originNode == target) {
                return clonedNode;
            }
            if (originNode->left) {
                originQ.push(originNode->left);
                clonedQ.push(clonedNode->left);
            }
            if (originNode->right) {
                originQ.push(originNode->right);
                clonedQ.push(clonedNode->right);
            }
        }
        return nullptr;
    }
};

int main(int argc, char const *argv[]) {
    return 0;
}
