#include <map>
#include <sstream>
#include <stack>
#include <string>
using namespace std;

class Solution {
   public:
    int calculate(string s) {
        stack<int> ops;
        ops.push(1);
        int sign = 1;

        int ret = 0;
        int n = s.length();
        int i = 0;
        while (i < n) {
            if (s[i] == ' ') {
                i++;
            } else if (s[i] == '+') {
                sign = ops.top();
                i++;
            } else if (s[i] == '-') {
                sign = -ops.top();
                i++;
            } else if (s[i] == '(') {
                ops.push(sign);
                i++;
            } else if (s[i] == ')') {
                ops.pop();
                i++;
            } else {
                long num = 0;
                while (i < n && s[i] >= '0' && s[i] <= '9') {
                    num = num * 10 + s[i] - '0';
                    i++;
                }
                ret += sign * num;
            }
        }
        return ret;
    }
};

/* 修改过后的四则计算器：支持表达式中存在负数、空格
class Solution {
   private:
    std::map<char, int> prior = {
        std::make_pair('+', 0),
        std::make_pair('-', 0),
        std::make_pair('*', 1),
        std::make_pair('/', 1),
    };
    std::string parseNum(const std::string &str) {
        int len = str.length();
        int pos;
        for (pos = 0; pos < len && isdigit(str[pos]); pos++)
            ;
        return str.substr(0, pos);
    }

    int calcPostfix(const std::string &postfix) {
        std::istringstream iss(postfix);
        std::stack<int> stack;
        for (int pos = 0; pos != std::string::npos;
             pos = postfix.find(" ", pos + 1)) {
            if (pos == 0 || isdigit(postfix[pos + 1])) {
                int num;
                iss >> num;
                stack.push(num);
            } else {
                int rhs = stack.top();
                stack.pop();
                int lhs = stack.top();
                stack.pop();
                char op;
                iss >> op;
                switch (op) {
                    case '+':
                        stack.push(lhs + rhs);
                        break;
                    case '-':
                        stack.push(lhs - rhs);
                        break;
                    case '*':
                        stack.push(lhs * rhs);
                        break;
                    case '/':
                        stack.push(lhs / rhs);
                        break;
                    default:
                        break;
                }
            }
        }
        return stack.top();
    }

    std::string infix2Postfix(const std::string &infix) {
        std::string postFix;
        std::stack<char> symStack;
        int len = infix.length();
        for (int i = 0; i < len; i++) {
            if (infix[i] == ' ') continue;
            if (isdigit(infix[i])) {
                std::string numStr = parseNum(infix.substr(i));
                if (postFix.empty()) {
                    postFix += numStr;
                } else {
                    postFix = postFix + " " + numStr;
                }
                i += numStr.length() - 1;
            } else if (infix[i] == '(') {
                symStack.push(infix[i]);
                if (infix[i + 1] == '-') postFix = postFix + " 0";
            } else if (infix[i] == ')') {
                while (symStack.top() != '(') {
                    postFix = postFix + " " + symStack.top();
                    symStack.pop();
                }
                symStack.pop();
            } else {
                while (!symStack.empty() && symStack.top() != '(' &&
                       prior[symStack.top()] >= prior[infix[i]]) {
                    postFix = postFix + " " + symStack.top();
                    symStack.pop();
                }
                symStack.push(infix[i]);
            }
        }
        while (!symStack.empty()) {
            postFix = postFix + " " + symStack.top();
            symStack.pop();
        }
        return postFix;
    }

   public:
    int calculate(std::string s) {
        if (s[0] == '-') s = "0" + s;
        std::string postStr = infix2Postfix(s);
        int ret = calcPostfix(postStr);
        return ret;
    }
};
*/
