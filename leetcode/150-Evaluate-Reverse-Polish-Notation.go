func evalRPN(tokens []string) int {
    var operands []int
    for _, token := range tokens {
        num, err := strconv.Atoi(token)
        if err == nil {
            operands = append(operands, num)
            continue
        }
        lhs, rhs := operands[len(operands)-2], operands[len(operands)-1]
        operands = operands[:len(operands) - 2]
        switch token {
        case "+":
            operands = append(operands, lhs + rhs)
        case "-":
            operands = append(operands, lhs - rhs)
        case "*":
            operands = append(operands, lhs * rhs)
        case "/":
            operands = append(operands, lhs / rhs)
        }
    }
    return operands[0]
}
