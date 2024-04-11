func countOfAtoms(formula string) string {
    idx := 0
    st := []map[string]int{{}}
    for idx < len(formula) {
        cnts := st[len(st) - 1]
        if formula[idx] == '(' {
            st = append(st, map[string]int{})
            idx++
        } else if formula[idx] == ')' {
            st = st[:len(st) - 1]
            idx++
            num, length := parseNum(formula[idx:])
            idx += length
            for atom, cnt := range cnts {
                st[len(st) - 1][atom] += cnt * num
            }
        } else {
            atom, length := parseAtom(formula[idx:])
            idx += length
            num, length := parseNum(formula[idx:])
            cnts[atom] += num
            idx += length
        }
    }

    cnts := st[0]
    type pair struct {
        atom string
        cnt int
    }
    slice := []pair{}
    for atom, cnt := range cnts {
        slice = append(slice, pair{atom, cnt})
    }
    sort.Slice(slice, func (lhs, rhs int) bool {
        return slice[lhs].atom < slice[rhs].atom
    })
    ret := []byte{}
    for _, elem := range slice {
        ret = append(ret, elem.atom...)
        if elem.cnt > 1 {
            ret = append(ret, strconv.Itoa(elem.cnt)...)
        }
    }
    return string(ret)
}

func parseAtom(str string) (string, int) {
    length := 1
    for length < len(str) && unicode.IsLower(rune(str[length])) {
        length++
    }
    return str[:length], length
}

func parseNum(str string) (int, int) {
    num, length := 0, 0
    for length < len(str) && unicode.IsDigit(rune(str[length])) {
        num, length = num * 10 + int(str[length] - '0'), length + 1
    }
    if num == 0 { return 1, 0 }
    return num, length
}
