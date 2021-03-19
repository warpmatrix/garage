func isValidSerialization(preorder string) bool {
    num := 1
    for i := 0; i < len(preorder); i++ {
        if preorder[i] == ',' {
            continue
        }
        if num <= 0 {
            return false
        }
        if preorder[i] == '#' {
            num--
        } else {
            num++
            for i < len(preorder) && preorder[i] != '#' {
                i++
            }
        }
    }
    return num == 0
}
