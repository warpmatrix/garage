func xorQueries(arr []int, queries [][]int) []int {
    xors := make([]int, len(arr) + 1)
    for i := range arr {
        xors[i + 1] = xors[i] ^ arr[i]
    }
    ret := make([]int, len(queries))
    for i := range queries {
        ret[i] = xors[queries[i][0]] ^ xors[queries[i][1] + 1]
    }
    return ret
}
