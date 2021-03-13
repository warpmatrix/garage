#include <string>
using std::string;
bool isValidSerialization(string preorder) {
    int num = 1;
    for (int i = 0; i < preorder.length(); i++) {
        if (preorder[i] == ',') continue;
        if (num <= 0) return false;
        if (isdigit(preorder[i])) {
            num++;
            while (i + 1 < preorder.length() && isdigit(preorder[i + 1])) i++;
        } else if (preorder[i] == '#') {
            num--;
        }
    }
    return num == 0;
}
