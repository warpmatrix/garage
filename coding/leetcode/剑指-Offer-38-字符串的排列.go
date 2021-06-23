func permutation(s string) []string {
    arr := []byte(s)
    sort.Slice(arr, func (lhs, rhs int) bool {
        return arr[lhs] < arr[rhs]
    })
    ret := []string{}
    for ; arr != nil; arr = nextPerm(arr) {
        ret = append(ret, string(arr))
    }
    return ret
}

func nextPerm(arr []byte) []byte {
    idx1 := len(arr) - 2
    for idx1 >= 0 && arr[idx1 + 1] <= arr[idx1] {
        idx1--
    }
    if idx1 == -1 {
        return nil
    }
    idx2 := len(arr) - 1
    for arr[idx2] <= arr[idx1] {
        idx2--
    }
    arr[idx1], arr[idx2] = arr[idx2], arr[idx1]
    reverse(arr[idx1 + 1:])
    return arr
}

func reverse(arr []byte) []byte {
    for i := 0; i * 2 < len(arr); i++ {
        tmp := len(arr) - i - 1
        arr[i], arr[tmp] = arr[tmp], arr[i]
    }
    return arr
}
