#include <iostream>
#include <map>
#include <stack>
#include <string>

std::map<char, int> prior = {
    std::make_pair('+', 0),
    std::make_pair('-', 0),
    std::make_pair('*', 1),
    std::make_pair('/', 1),
};
int parseNum(const std::string &str, int &len) {
    int num = 0;
    for (len = 0; len < str.length() && isdigit(str[len]); len++) {
        num = num * 10 + (str[len] - '0');
    }
    return num;
}
int calc(int lhs, int rhs, char op) {
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
            return 0;
    }
}
void reduceEquation(std::stack<int> &operands, std::stack<char> &operators) {
    int rhs = operands.top();
    operands.pop();
    int lhs = operands.top();
    operands.pop();
    char op = operators.top();
    operators.pop();
    int res = calc(lhs, rhs, op);
    operands.push(res);
}

int calculate(std::string s) {
    std::stack<char> operators;
    std::stack<int> operands;
    if (s[0] == '-') operands.push(0);
    for (int i = 0; i < s.length(); i++) {
        if (s[i] == ' ') continue;
        if (isdigit(s[i])) {
            int digitLen;
            int num = parseNum(s.substr(i), digitLen);
            i += digitLen - 1;
            operands.push(num);
        } else if (s[i] == '(') {
            operators.push(s[i]);
            if (i + 1 < s.length() && s[i + 1] == '-') operands.push(0);
        } else if (s[i] == ')') {
            while (operators.top() != '(') {
                reduceEquation(operands, operators);
            }
            // remove '('
            operators.pop();
        } else {
            while (!operators.empty() && operators.top() != '(' &&
                   prior[operators.top()] >= prior[s[i]]) {
                reduceEquation(operands, operators);
            }
            operators.push(s[i]);
        }
    }
    while (!operators.empty()) {
        reduceEquation(operands, operators);
    }
    return operands.top();
}

int main(int argc, char const *argv[]) {
    freopen("sample.in", "r", stdin);
    std::string str;
    std::cin >> str;
    std::cout << calculate(str) << '\n';
    return 0;
}
