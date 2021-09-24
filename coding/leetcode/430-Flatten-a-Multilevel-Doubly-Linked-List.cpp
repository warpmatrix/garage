/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* prev;
    Node* next;
    Node* child;
};
*/

class Solution {
public:
    Node* flatten(Node* head) {
        std::stack<Node *> st;
        Node *node = head;
        while (node != nullptr) {
            if (node->child) {
                if (node->next) {
                    st.push(node->next);
                }
                node->child->prev = node;
                node->next = node->child;
                node->child = nullptr;
            } else if (!node->next && !st.empty()) {
                node->next = st.top();
                node->next->prev = node;
                st.pop();
            }
            node = node->next;
        }
        return head;
    }
};
