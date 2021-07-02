func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)
    ret := 0
    for _, cost := range costs {
        if coins < cost { break }
        ret, coins = ret + 1, coins - cost
    }
    return ret
}
