#include <stack>
#include <string>
#include <sstream>
#include <vector>
using std::string;
using std::vector;

int reduceExp(char op, std::stack<int>& operands) {
    int rhs = operands.top();
    operands.pop();
    int lhs = operands.top();
    operands.pop();
    switch (op) {
        case '+':
            return lhs + rhs;
        case '-':
            return lhs - rhs;
        case '*':
            return lhs * rhs;
        case '/':
            return lhs / rhs;
        default:
            return -1;
    }
}

int str2int(string s) {
    std::stringstream ss;
    int n;
    ss << s;
    ss >> n;
    return n;
}

int evalRPN(vector<string>& tokens) {
    std::stack<int> operands;
    for (auto&& token : tokens) {
        if (token[0] >= '0' && token[0] <= '9' || token.length() > 1) {
            int num = str2int(token);
            operands.push(num);
            continue;
        }
        operands.push(reduceExp(token[0], operands));
    }
    return operands.top();
}
